package main

import (
	"wh-git.mingyuanyun.com/grpc-go/go-grpc-base/server"
	"wh-git.mingyuanyun.com/grpc-go/go-grpc-base/test/pb/demo"
	"context"
	"fmt"
)

var message *string

// DemoServer demo
type DemoServer struct {
}

// SayHello test
func (s *DemoServer) SayHello(ctx context.Context, who *demo.Who) (*demo.HelloEnvoy, error) {
	fmt.Printf("%s say hello to envoy\n", who.GetName())

	return &demo.HelloEnvoy{
		Message: fmt.Sprintf("Wow!!!!%s %s", *message, who.GetName()),
	}, nil
}

func main() {
	fmt.Printf("This is test!")

	s := server.NewServer(server.WithPort(8888))
	demo.RegisterHelloServer(s.GrpcServer(), &DemoServer{})
	_ = s.StartServer()

	//lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8888))
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	///**
	//	在 gRPC 中，大类可分为两种 RPC 方法，与拦截器的对应关系是：
	//		普通方法：一元拦截器（grpc.UnaryInterceptor）
	//		流方法：流拦截器（grpc.StreamInterceptor）
	//	grpc_middleware: 		将多个拦截器组合成一个拦截器链
	//	StreamInterceptor: 		对于stream调用的拦截器
	//	UnaryInterceptor: 		对于服单向调用的拦截器
	//	grpc_ctxtags: 			为上下文增加 Tag, map对象
	//	grpc_opentracing:		支持opentracing/zipkin
	//	grpc_zap:				zap日志框架
	//	grpc_auth:				身份验证拦截器
	//	grpc_recovery: 			cu
	//*/
	//grpcServer := grpc.NewServer(
	//	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	//		grpc_ctxtags.StreamServerInterceptor(),
	//		grpc_opentracing.StreamServerInterceptor(),
	//		grpc_recovery.StreamServerInterceptor(),
	//	)),
	//	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
	//		grpc_ctxtags.UnaryServerInterceptor(),
	//		grpc_opentracing.UnaryServerInterceptor(),
	//		grpc_recovery.UnaryServerInterceptor(),
	//	)),
	//)
	//svr := &DemoServer{}
	//demo.RegisterHelloServer(grpcServer, svr)
	//fmt.Println("hello demo service ready")
	//_ = grpcServer.Serve(lis)
}
