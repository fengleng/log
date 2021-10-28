package log

import (
	"github.com/fengleng/go-common/fileutil"
	"os"
)

func NewFileLogger(opts ...CfgOption) Logger {
	cfg := defaultLogCfg

	for _, f := range opts {
		f(cfg)
	}
	cfg.logOutputType = OutputTypeFile

	if !fileutil.FileExists(cfg.dir) {
		err := os.MkdirAll(cfg.dir, 0666)
		panic(err)
	}

	return newXLog(cfg)
}
