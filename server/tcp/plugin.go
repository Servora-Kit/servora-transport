package tcp

import (
	"context"
	"fmt"

	tcpconf "github.com/Servora-Kit/servora-transport/server/tcp/gen/conf"
	"github.com/Servora-Kit/servora/transport/runtime"
)

const (
	Type = "tcp"

	// ExtraKeyHandler allows runtime graph to inject a connection handler.
	ExtraKeyHandler = "handler"
)

// Plugin adapts TCP server to Servora transport runtime graph.
type Plugin struct{}

func (p *Plugin) Type() string { return Type }

func (p *Plugin) Build(_ context.Context, in runtime.ServerBuildInput) (runtime.Server, error) {
	opts := make([]ServerOption, 0, 3)

	if in.Config != nil {
		cfg, ok := in.Config.(*tcpconf.Server)
		if !ok {
			return nil, fmt.Errorf("tcp plugin expects *tcpconf.Server config, got %T", in.Config)
		}
		opts = append(opts, WithConfig(cfg))
	}
	if in.Logger != nil {
		opts = append(opts, WithLogger(in.Logger))
	}
	if len(in.ExtraValues) > 0 {
		if raw, ok := in.ExtraValues[ExtraKeyHandler]; ok && raw != nil {
			h, ok := raw.(ConnectionHandler)
			if !ok {
				return nil, fmt.Errorf("tcp plugin expects ConnectionHandler for %q, got %T", ExtraKeyHandler, raw)
			}
			opts = append(opts, WithConnectionHandler(h))
		}
	}

	return NewServer(opts...), nil
}
