package log

import (
	"github.com/fengleng/go-common/file"
	"path/filepath"
)

func NewFileLogger(opts ...CfgOption) Logger {
	cfg := defaultLogCfg

	for _, f := range opts {
		f(cfg)
	}
	cfg.logOutputType = OutputTypeFile

	return newXLog(cfg)
}
