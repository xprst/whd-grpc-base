package req_uuid

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
)

type requestIDKey struct {
}


func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestID := HandleRequestID(ctx)
		ctx = context.WithValue(ctx, requestIDKey{}, requestID)
		grpc_ctxtags.Extract(ctx).Set("grpc.request.id", requestID)
		return handler(ctx, req)
	}
}

type serverStreamWithContext struct {
	grpc.ServerStream
	ctx context.Context
}

func (ss serverStreamWithContext) Context() context.Context {
	return ss.ctx
}

func NewServerStreamWithContext(stream grpc.ServerStream, ctx context.Context) grpc.ServerStream {
	return serverStreamWithContext{
		ServerStream: stream,
		ctx:          ctx,
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {

	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		ctx := stream.Context()
		requestID := HandleRequestID(ctx)
		ctx = context.WithValue(ctx, requestIDKey{}, requestID)
		grpc_ctxtags.Extract(ctx).Set("grpc.request.id", requestID)
		stream = NewServerStreamWithContext(stream, ctx)
		return handler(srv, stream)
	}
}

func FromContext(ctx context.Context) string {
	id, ok := ctx.Value(requestIDKey{}).(string)
	if !ok {
		return ""
	}
	return id
}
