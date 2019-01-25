package app

import (
	"github.com/xprst/whd-grpc-base/config"
	"github.com/xprst/whd-grpc-base/orm"
	"go.uber.org/zap"
)

// 全局日志
var Logger *zap.Logger
// 全局配置项
var Config config.ConfigEngine
// 全局数据连接
var MyBean *orm.MyBean

