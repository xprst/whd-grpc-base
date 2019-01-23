package log_zap

/**
	日志中间件
 */
import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"os"
)

// LoggingInterceptor RPC 方法的入参出参的日志输出
//func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
//	resp, err := handler(ctx, req)
//	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
//	return resp, err
//}

var logger *zap.Logger

func initLogger(logPath string) *zap.Logger {
	// 仅打印Error级别以上的日志
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	// 打印所有级别的日志
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	// 控制台输出
	zapConsole := zapcore.Lock(os.Stdout)
	// 日志文件输出
	zapWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    1023,    // megabytes
		MaxBackups: 3,       // 最多保留3个备份
		MaxAge:     14,      // days
		Compress:   false,   // 是否压缩
	})

	//encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapConsole, lowPriority), // 打印在控制台
		zapcore.NewCore(consoleEncoder, zapWriter, highPriority), // 打印在文件中
	)

	opts := []zap.Option{
		zap.ErrorOutput(zapWriter),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	}
	logger := zap.New(core, opts...)
	logger.Info("DefaultLogger init success")
	grpc_zap.ReplaceGrpcLogger(logger)

	return logger

}

func WithZapLogger() *zap.Logger {
	if logger == nil {
		logger = initLogger("./logs/request.log")
	}

	return logger
}

func UnaryServerInterceptor(zapLogger *zap.Logger) grpc.UnaryServerInterceptor {
	return grpc_zap.UnaryServerInterceptor(zapLogger)
}

func StreamServerInterceptor(zapLogger *zap.Logger) grpc.StreamServerInterceptor {
	return grpc_zap.StreamServerInterceptor(zapLogger)
}
