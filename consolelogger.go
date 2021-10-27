package log

import (
	"gitee.com/fenleng/flyfisher/model"
)

func NewConsoleLogger(m *model.LogCfg) Logger {
	cfg := make(map[string]string)
	if m == nil {
		m = &model.LogCfg{}
	}
	if m.Level != "" {
		cfg["level"] = m.Level
	}

	if m.Service != "" {
		cfg["service"] = m.Service
	}
	return NewXLog(cfg)
}
