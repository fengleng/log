package log

func NewFileLogger(opts ...CfgOption) Logger {
	cfg := defaultLogCfg

	for _, f := range opts {
		f(cfg)
	}
	cfg.logOutputType = OutputTypeFile

	return newXLog(cfg)
}
