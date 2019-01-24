package server

import (
	"errors"
	"fmt"
	"github.com/xprst/whd-grpc-base/middleware"
	"golang.org/x/net/trace"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"net"
)

var ErrServerClosed = errors.New("grpc: server closed")

// Server grpc server
type Server struct {
	port       int32
	Plugins    PluginContainer
	options    map[string]interface{}
	grpcServer *grpc.Server
}

// NewServer 通过option函数修改Server的属性
func NewServer(options ...OptionFn) *Server {

	s := &Server{
		Plugins:    &pluginContainer{},
		options:    make(map[string]interface{}),
		grpcServer: grpc.NewServer(middleware.NewOptions()...),
	}

	for _, op := range options {
		op(s)
	}

	return s
}

func (s *Server) GrpcServer() *grpc.Server {
	return s.grpcServer
}

// RegisterService 注册grpc服务
func (s *Server) RegisterService() {

}

// StartServer 启动服务
func (s *Server) StartServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return ErrServerClosed
	}
	go startTrace()
	fmt.Println("grpc service start at port:", s.port)
	return s.grpcServer.Serve(lis)
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	fmt.Println("Trace listen on 50051")
}
