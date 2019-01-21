package log

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v)
}

// LoggingInterceptor RPC 方法的入参出参的日志输出
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}