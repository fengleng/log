package log

import (
	"gitee.com/fenleng/flyfisher/model"
)

func NewFileLogger(m *model.LogCfg) Logger {
	if m == nil {
		m = &model.LogCfg{}
	}
	cfg := make(map[string]string)
	if m.Level != "" {
		cfg["level"] = m.Level
	}
	if m.Dir != "" {
		cfg["dir"] = m.Dir
	}
	if m.LogFileName != "" {
		cfg["logFileName"] = m.LogFileName
	}
	if m.Service != "" {
		cfg["service"] = m.Service
	}
	if m.IsLogFile {
		cfg["isLogFile"] = "true"
	}
	return NewXLog(cfg)
}
