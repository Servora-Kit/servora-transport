package tcp

import (
	"context"
	"testing"
	"time"

	tcpconf "github.com/Servora-Kit/servora-transport/server/tcp/gen/conf"
	confv1 "github.com/Servora-Kit/servora/api/gen/go/servora/conf/v1"
)

func TestServerEndpointAndLifecycle(t *testing.T) {
	srv := NewServer(
		WithConfig(&tcpconf.Server{Listen: &confv1.Server_Listen{Addr: ":0"}}),
	)

	ep, err := srv.Endpoint()
	if err != nil {
		t.Fatalf("endpoint: %v", err)
	}
	if ep == nil || ep.Host == "" {
		t.Fatalf("unexpected endpoint: %#v", ep)
	}

	if err := srv.Start(context.Background()); err != nil {
		t.Fatalf("start: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Stop(ctx); err != nil {
		t.Fatalf("stop: %v", err)
	}
}

func TestServerRespectsExplicitRegistryEndpoint(t *testing.T) {
	srv := NewServer(
		WithConfig(&tcpconf.Server{
			Listen:   &confv1.Server_Listen{Addr: ":0"},
			Registry: &confv1.Server_Registry{Endpoint: "tcp://127.0.0.1:9000"},
		}),
	)

	ep, err := srv.Endpoint()
	if err != nil {
		t.Fatalf("endpoint: %v", err)
	}
	if got := ep.String(); got != "tcp://127.0.0.1:9000" {
		t.Fatalf("endpoint = %s, want %s", got, "tcp://127.0.0.1:9000")
	}
}
