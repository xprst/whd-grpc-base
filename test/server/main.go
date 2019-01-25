package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/xprst/whd-grpc-base/server"
	"github.com/xprst/whd-grpc-base/test/pb/demo"
	"google.golang.org/grpc/metadata"
)

var message *string

// DemoServer demo
type DemoServer struct {
}

// SayHello test
func (s *DemoServer) SayHello(ctx context.Context, who *demo.Who) (*demo.HelloEnvoy, error) {
	fmt.Printf("%s say hello to envoy\n", who.GetName())

	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md)

	v := grpc_ctxtags.Extract(ctx).Values()
	fmt.Println(v)

	name := who.GetName()
	if name == "panic" {
		panic("test error ...")
	}
	if name == "" {
		return nil, errors.New("param name is need, but name is null!")
	}

	return &demo.HelloEnvoy{
		Message: fmt.Sprintf("Wow!!!!%s %s", *message, who.GetName()),
	}, nil
}

func main() {
	fmt.Printf("This is test!")
	message = flag.String("m", "hello", "name aa")
	configPath := flag.String("conf","./config/app.json", "config file path")
	flag.Parse()

	s := server.NewServer(server.WithConfig(*configPath))
	demo.RegisterHelloServer(s.GrpcServer(), &DemoServer{})
	err := s.StartServer()
	if err != nil {
		panic(err)
	}
}
