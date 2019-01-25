package app

import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/xprst/whd-grpc-base/config"
	"github.com/xprst/whd-grpc-base/middleware/log"
	"github.com/xprst/whd-grpc-base/orm"
	"go.uber.org/zap"
)

// Grpc端口
var GrpcPort int32
// 全局日志
var Logger *zap.Logger
// 全局配置项
var Config config.ConfigEngine
// 全局数据连接
var MyBean *orm.MyBean

// InitWithConfig 从配置文件初始化应用程序
func InitWithConfig(path string) error {
	// 设置默认端口
	GrpcPort = 8080
	// 初始化日志
	Logger = log_zap.WithZapLogger()
	// 载入配置项
	err := Config.Load(path)
	if err != nil {
		return err
	}
	// 初始化orm容器
	MyBean = orm.NewBean()

	// 从配置文件读取grpc port
	GrpcPort = int32(Config.GetInt("server.grpcPort"))

	// 初始化数据连接
	var dataSources = make(map[string]orm.DataSource)
	Config.GetStruct("dataSource", dataSources)

	for k, ds := range dataSources {
		e, err := xorm.NewEngine(ds.DriverName, ds.DataSourceName)
		if err != nil {
			return err
		}
		ormLogLevel := orm.GetLevel(ds.LogLevel)
		ormLogger := orm.NewXormLogger(Logger, ormLogLevel)
		if ormLogLevel == core.LOG_DEBUG{
			ormLogger.ShowSQL(true)
		}
		e.SetLogger(ormLogger)

		MyBean.RegisterEngine(k, e)
	}

	return nil
}
