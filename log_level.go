package log

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func (l Level) String() string {
	levelStr := "INFO"
	switch l {
	case DEBUG:
		levelStr = "DBG"
	case INFO:
		levelStr = "INF"
	case WARN:
		levelStr = "WAR"
	case ERROR:
		levelStr = "ERR"
	case FATAL:
		levelStr = "FAT"
	}
	return levelStr
}
