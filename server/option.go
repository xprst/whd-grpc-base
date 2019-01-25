package server

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/xprst/whd-grpc-base/app"
	"github.com/xprst/whd-grpc-base/middleware/log"
	"github.com/xprst/whd-grpc-base/orm"
)

// OptionFn configures options of server.
type OptionFn func(*Server)


// WithConfig 设置配置文件
func WithConfig(path string) OptionFn {
	return func(s *Server) {
		// 初始化日志
		app.Logger = log_zap.WithZapLogger()
		// 载入配置项
		err := app.Config.Load(path)
		if err != nil {
			panic(fmt.Sprintf("InitWithConfig failed ： %v",err))
		}
		// 初始化orm容器
		app.MyBean = orm.NewBean()

		// 初始化数据连接
		var dataSources = make(map[string]orm.DataSource)
		app.Config.GetStruct("dataSource", dataSources)

		for k, ds := range dataSources {
			e, err := xorm.NewEngine(ds.DriverName, ds.DataSourceName)
			if err != nil {
				panic(fmt.Sprintf("Init db conn failed. DB key：%s, error %v",k, err))
			}
			ormLogLevel := orm.GetLevel(ds.LogLevel)
			ormLogger := orm.NewXormLogger(app.Logger, ormLogLevel)
			if ormLogLevel == core.LOG_DEBUG {
				ormLogger.ShowSQL(true)
			}
			e.SetLogger(ormLogger)

			app.MyBean.RegisterEngine(k, e)
		}


		// 从配置文件读取grpc port
		s.port = int32(app.Config.GetInt("server.grpcPort"))
		// 设置默认端口
		if s.port == 0 {
			s.port = 8888
		}
	}
}
