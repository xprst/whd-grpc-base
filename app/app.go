package app

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xprst/whd-grpc-base/config"
	"github.com/xprst/whd-grpc-base/orm"
)

// Grpc端口
var GrpcPort int32
// 全局配置项
var Config config.ConfigEngine
// 全局数据连接
var MyBean *orm.MyBean

// InitWithConfig 从配置文件初始化应用程序
func InitWithConfig(path string) {
	GrpcPort = 8080
	MyBean = orm.NewBean()
	err := Config.Load(path)
	if err != nil {
		fmt.Sprintf("%v \n", err)
	}
	// 从配置文件读取grpc port
	GrpcPort = int32(Config.GetInt("server.grpcPort"))
	// 初始化数据连接
	var dataSources = make(map[string]config.DataSource)
	Config.GetStruct("dataSource", dataSources)

	for k, ds := range dataSources {
		e, err := xorm.NewEngine(ds.DriverName, ds.DataSourceName)
		if err != nil {
			fmt.Errorf("failed to create xorm engine for conn: %s . Error :%v",k, err)
		}
		MyBean.RegisterEngine(k, e)
	}
}
