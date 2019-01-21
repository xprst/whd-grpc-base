package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"runtime/debug"

	"github.com/xprst/whd-grpc-base/middleware/log"
	"google.golang.org/grpc"
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
	/**
		经验证发现：拦截器的用法，酷似koa的洋葱圈，有hangler方法 ==> await next（）
		拦截器特点：后面的先生效，且覆盖前面的拦截器
	 */
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			RecoveryInterceptor,	// 该拦截器一定放第一个，已保护异常处理
			log.LoggingInterceptor,
		),
	}
	s := &Server{
		Plugins:    &pluginContainer{},
		options:    make(map[string]interface{}),
		grpcServer: grpc.NewServer(opts...),
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
	fmt.Println("grpc service ready")
	return s.grpcServer.Serve(lis)
}

// RecoveryInterceptor RPC 方法的异常保护和日志输出
func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()

	return handler(ctx, req)
}
