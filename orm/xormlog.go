package orm

import (
	"fmt"
	"github.com/go-xorm/core"
	"go.uber.org/zap"
)

// logger interface
// XormLogger is the default implment of core.ILogger
type XormLogger struct {
	logger *zap.Logger
	level   core.LogLevel
	showSQL bool
}

// NewXormLogger use a special io.Writer as logger output
func NewXormLogger(log *zap.Logger, l core.LogLevel) *XormLogger {
	return &XormLogger{
		logger: log,
		level: l,
	}
}

// Error implement core.ILogger
func (s *XormLogger) Error(v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.logger.Error(fmt.Sprintf("[xorm] %v", v...))
	}
	return
}

// Errorf implement core.ILogger
func (s *XormLogger) Errorf(format string, v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.logger.Error(fmt.Sprintf("[xorm] %s",fmt.Sprintf(format, v...)))
	}
	return
}

// Debug implement core.ILogger
func (s *XormLogger) Debug(v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.logger.Debug(fmt.Sprintf("[xorm] %v", v...))
	}
	return
}

// Debugf implement core.ILogger
func (s *XormLogger) Debugf(format string, v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.logger.Debug(fmt.Sprintf("[xorm] %s",fmt.Sprintf(format, v...)))
	}
	return
}

// Info implement core.ILogger
func (s *XormLogger) Info(v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.logger.Info(fmt.Sprintf("[xorm] %v", v...))
	}
	return
}

// Infof implement core.ILogger
func (s *XormLogger) Infof(format string, v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.logger.Info(fmt.Sprintf("[xorm] %s",fmt.Sprintf(format, v...)))
	}
	return
}

// Warn implement core.ILogger
func (s *XormLogger) Warn(v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.logger.Warn(fmt.Sprintf("[xorm] %v", v...))
	}
	return
}

// Warnf implement core.ILogger
func (s *XormLogger) Warnf(format string, v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.logger.Warn(fmt.Sprintf("[xorm] %s",fmt.Sprintf(format, v...)))
	}
	return
}

// Level implement core.ILogger
func (s *XormLogger) Level() core.LogLevel {
	return s.level
}

// SetLevel implement core.ILogger
func (s *XormLogger) SetLevel(l core.LogLevel) {
	s.level = l
	return
}

// ShowSQL implement core.ILogger
func (s *XormLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

// IsShowSQL implement core.ILogger
func (s *XormLogger) IsShowSQL() bool {
	return s.showSQL
}