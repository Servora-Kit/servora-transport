package tcp

import (
	"context"
	"net"

	tcpconf "github.com/Servora-Kit/servora-transport/server/tcp/gen/conf"
	"github.com/go-kratos/kratos/v2/log"
)

// ConnectionHandler handles accepted TCP connections.
type ConnectionHandler func(ctx context.Context, conn net.Conn)

type ServerOption func(*serverOptions)

type serverOptions struct {
	config  *tcpconf.Server
	logger  log.Logger
	handler ConnectionHandler
}

func WithConfig(c *tcpconf.Server) ServerOption {
	return func(o *serverOptions) {
		o.config = c
	}
}

func WithLogger(l log.Logger) ServerOption {
	return func(o *serverOptions) {
		o.logger = l
	}
}

func WithConnectionHandler(h ConnectionHandler) ServerOption {
	return func(o *serverOptions) {
		o.handler = h
	}
}
