package log

import (
	"github.com/fengleng/go-common/fileutil"
)

func NewFileLogger(opts ...CfgOption) Logger {
	cfg := defaultLogCfg

	for _, f := range opts {
		f(cfg)
	}
	cfg.logOutputType = OutputTypeFile

	utils.FileExist()

	return newXLog(cfg)
}
