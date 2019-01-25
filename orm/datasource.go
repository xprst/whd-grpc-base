package orm

import "github.com/go-xorm/core"


type DataSource struct {
	DriverName string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
	LogLevel string `json:"logLevel"`
}

// GetLevel 解析字符串level到core.LogLevel
func GetLevel(logLevel string) core.LogLevel {
	switch logLevel {
	case "debug":
		return core.LOG_DEBUG
	case "info":
		return core.LOG_INFO
	case "warn":
		return core.LOG_WARNING
	case "error":
		return core.LOG_ERR
	default:
		return core.LOG_OFF
	}
}
