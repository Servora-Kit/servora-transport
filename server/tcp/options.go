package tcp

import (
	"context"
	"log/slog"
	"net"

	tcpconf "github.com/Servora-Kit/servora-transport/server/tcp/gen/conf"
)

const Type = "tcp"

// ConnectionHandler handles accepted TCP connections.
type ConnectionHandler func(ctx context.Context, conn net.Conn)

type ServerOption func(*serverOptions)

type serverOptions struct {
	config  *tcpconf.Server
	logger  *slog.Logger
	handler ConnectionHandler
}

func WithConfig(c *tcpconf.Server) ServerOption {
	return func(o *serverOptions) {
		o.config = c
	}
}

func WithLogger(l *slog.Logger) ServerOption {
	return func(o *serverOptions) {
		o.logger = l
	}
}

func WithConnectionHandler(h ConnectionHandler) ServerOption {
	return func(o *serverOptions) {
		o.handler = h
	}
}
