package log

import (
	"github.com/fengleng/go-common/file_cfg"
	"io"
	"os"
)

type OutputType int

const (
	OutputTypeStd   = iota + 1
	OutputTypeFile  = iota + 1
	OutputTypeQueue = iota + 1
)

type logCfg struct {
	logOutputType OutputType
	logLevel      Level

	wc             io.WriteCloser
	logChannelSize int

	dir        string
	filePrefix string

	service string
}

var defaultLogCfg = &logCfg{
	logOutputType:  OutputTypeStd,
	wc:             os.Stdout,
	dir:            file_cfg.LogDir,
	filePrefix:     file_cfg.LogPreFixName,
	service:        "XLOG",
	logLevel:       INFO,
	logChannelSize: 50000,
}

type CfgOption func(opt *logCfg)

func OutPutCfgOption(t OutputType) CfgOption {
	return func(opt *logCfg) {
		opt.logOutputType = t
	}
}

func WCCfgOption(wc io.WriteCloser) CfgOption {
	return func(opt *logCfg) {
		opt.wc = wc
	}
}

func DirCfgOption(dir string) CfgOption {
	return func(opt *logCfg) {
		opt.dir = dir
	}
}

func FilePrefixCfgOption(filePrefix string) CfgOption {
	return func(opt *logCfg) {
		opt.filePrefix = filePrefix
	}
}

func ServiceCfgOption(service string) CfgOption {
	return func(opt *logCfg) {
		opt.service = service
	}
}

func ChannelSizeCfgOption(logChannelSize int) CfgOption {
	return func(opt *logCfg) {
		opt.logChannelSize = logChannelSize
	}
}
