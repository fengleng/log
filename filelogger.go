package log

import (
	"github.com/fengleng/go-common/fl_file"
)

func NewFileLogger(opts ...CfgOption) Logger {
	cfg := defaultLogCfg

	for _, f := range opts {
		f(cfg)
	}
	cfg.logOutputType = OutputTypeFile

	//fl_file.

	return newXLog(cfg)
}
