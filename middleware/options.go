package middleware

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/xprst/whd-grpc-base/middleware/log"
	"github.com/xprst/whd-grpc-base/middleware/reqid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"
)

func NewOptions() []grpc.ServerOption {

	/**
		经验证发现：拦截器的用法，酷似koa的洋葱圈，有hangler方法 ==> await next（）
		拦截器特点：后面的先生效，且覆盖前面的拦截器
	 */
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(handlerRecovery)), // 该拦截器一定放第一个，已保护异常处理
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			log_zap.UnaryServerInterceptor(log_zap.WithZapLogger()),
			req_uuid.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(handlerRecovery)),
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			log_zap.StreamServerInterceptor(log_zap.WithZapLogger()),
			req_uuid.StreamServerInterceptor(),
		),
	}

	return opts
}

func handlerRecovery(p interface{}) (err error) {
	log_zap.WithZapLogger().Error(string(debug.Stack()[:]))
	return status.Errorf(codes.Internal, "%s", p)
}
