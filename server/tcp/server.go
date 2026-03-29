package tcp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"sync"
	"time"

	tcpconf "github.com/Servora-Kit/servora-transport/server/tcp/gen/conf"
	"github.com/Servora-Kit/servora/transport/shared/endpoint"
	"github.com/go-kratos/kratos/v2/log"
)

var errServerNotStarted = errors.New("tcp server not started")

// Server implements runtime.Server for TCP protocol.
type Server struct {
	mu      sync.Mutex
	opts    serverOptions
	lis     net.Listener
	ep      *url.URL
	closed  bool
	started bool
	wg      sync.WaitGroup
}

func NewServer(opts ...ServerOption) *Server {
	o := serverOptions{}
	for _, opt := range opts {
		if opt != nil {
			opt(&o)
		}
	}
	if o.handler == nil {
		o.handler = func(_ context.Context, conn net.Conn) {
			_ = conn.Close()
		}
	}
	return &Server{opts: o}
}

func (s *Server) Start(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return errors.New("tcp server already stopped")
	}
	if s.started {
		return nil
	}
	if err := s.ensureListenerLocked(); err != nil {
		return err
	}

	s.started = true
	s.wg.Add(1)
	go s.acceptLoop(ctx)
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return nil
	}
	s.closed = true
	lis := s.lis
	s.lis = nil
	s.mu.Unlock()

	if lis != nil {
		_ = lis.Close()
	}

	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	if ctx == nil {
		<-done
		return nil
	}

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Server) Endpoint() (*url.URL, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.ensureListenerLocked(); err != nil {
		return nil, err
	}
	if s.ep == nil {
		return nil, errServerNotStarted
	}
	return s.ep, nil
}

func (s *Server) acceptLoop(ctx context.Context) {
	defer s.wg.Done()

	for {
		s.mu.Lock()
		lis := s.lis
		s.mu.Unlock()
		if lis == nil {
			return
		}

		conn, err := lis.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				time.Sleep(50 * time.Millisecond)
				continue
			}
			if s.opts.logger != nil {
				_ = s.opts.logger.Log(log.LevelError, "msg", "tcp accept failed", "err", err)
			}
			return
		}

		h := s.opts.handler
		s.wg.Add(1)
		go func(c net.Conn) {
			defer s.wg.Done()
			h(ctx, c)
		}(conn)
	}
}

func (s *Server) ensureListenerLocked() error {
	if s.lis != nil {
		return nil
	}

	network := "tcp"
	addr := ":0"
	cfg := s.opts.config
	if cfg != nil && cfg.GetListen() != nil {
		if v := cfg.GetListen().GetNetwork(); v != "" {
			network = v
		}
		if v := cfg.GetListen().GetAddr(); v != "" {
			addr = v
		}
	}

	lis, err := net.Listen(network, addr)
	if err != nil {
		return fmt.Errorf("listen tcp server: %w", err)
	}
	s.lis = lis

	bindAddr := lis.Addr().String()
	secure := cfg != nil && cfg.GetTls() != nil && cfg.GetTls().GetEnable()
	regEndpoint := ""
	regHost := ""
	if cfg != nil && cfg.GetRegistry() != nil {
		regEndpoint = cfg.GetRegistry().GetEndpoint()
		regHost = cfg.GetRegistry().GetHost()
	}

	ep, err := endpoint.ResolveRegistryEndpoint(Type, bindAddr, regEndpoint, regHost, secure)
	if err != nil {
		_ = lis.Close()
		s.lis = nil
		return fmt.Errorf("resolve tcp registry endpoint: %w", err)
	}
	if ep == nil {
		scheme := Type
		if secure {
			scheme = "tcps"
		}
		ep = &url.URL{Scheme: scheme, Host: bindAddr}
	}
	s.ep = ep
	return nil
}

func (s *Server) config() *tcpconf.Server {
	return s.opts.config
}
