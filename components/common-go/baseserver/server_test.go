package baseserver_test

import (
	"context"
	"fmt"
	"github.com/gitpod-io/gitpod/common-go/baseserver"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	s := baseserver.New("test", baseserver.Opts{
		GRPCPort: 9001,
		HTTPPort: 9000,
	})

	go func() {
		if err := s.ListenAndServe(); err != nil {
			t.Fatal("Server failed to start")
		}
	}()

	time.Sleep(5 * time.Second)

	if err := s.Close(context.Background()); err != nil {
		t.Fatal("failed to shut down server")
	}

	fmt.Println("server terminated")
}

func TestServer2(t *testing.T) {
	s := baseserver.New("test", baseserver.Opts{
		GRPCPort: 9001,
		HTTPPort: 9000,
	})

	go func() {
		if err := s.ListenAndServe(); err != nil {
			t.Fatal("Server failed to start")
		}
	}()

	time.Sleep(5 * time.Second)

	if err := s.Close(context.Background()); err != nil {
		t.Fatal("failed to shut down server")
	}

	fmt.Println("server terminated")
}