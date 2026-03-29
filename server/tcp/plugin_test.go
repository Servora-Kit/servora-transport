package tcp

import (
	"context"
	"testing"

	tcpconf "github.com/Servora-Kit/servora-transport/server/tcp/gen/conf"
	confv1 "github.com/Servora-Kit/servora/api/gen/go/servora/conf/v1"
	"github.com/Servora-Kit/servora/transport/runtime"
)

func TestPluginType(t *testing.T) {
	if (&Plugin{}).Type() != Type {
		t.Fatalf("unexpected plugin type")
	}
}

func TestPluginBuild(t *testing.T) {
	srv, err := (&Plugin{}).Build(context.Background(), runtime.ServerBuildInput{
		Config: &tcpconf.Server{Listen: &confv1.Server_Listen{Addr: ":0"}},
	})
	if err != nil {
		t.Fatalf("build plugin: %v", err)
	}
	if srv == nil {
		t.Fatal("expected non-nil server")
	}
}

func TestPluginBuildRejectsWrongConfigType(t *testing.T) {
	_, err := (&Plugin{}).Build(context.Background(), runtime.ServerBuildInput{Config: "invalid"})
	if err == nil {
		t.Fatal("expected config type error")
	}
}
