package log

import "gitee.com/fenleng/flyfisher/model"

type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

var (
	logger Logger
)

func init() {
	logger = NewConsoleLogger(nil)
}

func InitLogger(m *model.LogCfg) {
	if m == nil {
		m = &model.LogCfg{}
	}
	var lg Logger
	if m.IsLogFile {
		lg = NewFileLogger(m)
	} else {
		lg = NewConsoleLogger(m)
	}
	logger = lg
}

func Debug(format string, a ...interface{}) {
	logger.Debug(format, a...)
}

func Error(format string, a ...interface{}) {
	logger.Error(format, a...)
}

func Info(format string, a ...interface{}) {
	logger.Info(format, a...)
}

func Warn(format string, a ...interface{}) {
	logger.Warn(format, a...)
}

func Fatal(format string, a ...interface{}) {
	logger.Fatal(format, a...)
}
