package log

import (
	"bytes"
	"fmt"
	"gitee.com/fenleng/flyfisher/core/hack"
	log2 "log"
	"runtime"
	"sync"

	"io"
	"os"
	"path"
	"time"
)

const (
	defaultSkip = 4
)

var (
	defaultService     = "XLog"
	defaultLevel       = levelTextArray[DebugLevel]
	defaultLogDir      string
	defaultSep         string
	defaultLogFileName = "xlog"
)

func init() {
	if runtime.GOOS == "windows" {
		defaultSep = `\`
		defaultLogDir = `c:\fl\xlog`
	} else {
		defaultSep = `/`
		defaultLogDir = `/var/log/fl/xlog`
	}
}

type XLog struct {
	level int

	skip     int
	hostname string
	service  string

	logTextChan chan string
	isLogFile   bool
	w           io.WriteCloser
	dir         string
	sep         string
	fileName    string
	curHour     int
	curHourTime string

	openLock sync.Mutex
}

func NewXLog(cfg map[string]string) *XLog {
	if cfg == nil {
		cfg = make(map[string]string)
	}
	_, ok := cfg["service"]
	if !ok {
		cfg["service"] = defaultService
	}
	_, ok = cfg["level"]
	if !ok {
		cfg["level"] = defaultLevel
	}
	level := LevelFromStr(cfg["level"])
	hostname, _ := os.Hostname()

	_, isLogFile := cfg["isLogFile"]
	if isLogFile {
		_, ok = cfg["dir"]
		if !ok {
			cfg["dir"] = defaultLogDir
		}
		_, ok = cfg["logFileName"]
		if !ok {
			cfg["logFileName"] = defaultLogFileName
		}
	}
	var (
		wc = os.Stdout
	)
	x := &XLog{
		level:       level,
		skip:        defaultSkip,
		service:     cfg["service"],
		hostname:    hostname,
		logTextChan: make(chan string, 5000),
		isLogFile:   isLogFile,
		dir:         cfg["dir"],
		fileName:    cfg["logFileName"],
		sep:         defaultSep,
		//curHour: time.Now().Hour(),
		//curHourTime: time.Now().Format("2006010215"),
		w: wc,
	}
	if x.isLogFile {
		x.reopen()
	}
	go x.Flush()
	return x
}

func (x *XLog) Flush() {
	for {
		select {
		case logText := <-x.logTextChan:
			x.write(logText)
		}
	}
}

func (x *XLog) write(logText string) {
	if x.needReopen() {
		x.reopen()
	}
	x.w.Write(hack.Slice(logText))
}

func (x *XLog) needReopen() bool {
	return x.isLogFile && x.curHour != time.Now().Hour()
}

func (x *XLog) reopen() {
	x.openLock.Lock()
	defer x.openLock.Unlock()
	if !x.needReopen() {
		return
	}
	x.curHour = time.Now().Hour()
	x.curHourTime = time.Now().Format("200601020315")
	fName := fmt.Sprintf("%s%s%s%s", x.dir, x.sep, x.fileName, x.curHourTime)
	file, err := os.OpenFile(fName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log2.Println(err)
		file, err = os.OpenFile(fName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			log2.Println(err)
			panic(err.Error())
		}
	}
	x.w = file
}

func (x *XLog) Debug(txt string, a ...interface{}) {
	if x.level > DebugLevel {
		return
	}
	x.formatLogAndWriteChan(formatValue(txt, a...), DebugLevel)
}

func (x *XLog) Info(txt string, a ...interface{}) {
	if x.level > InfoLevel {
		return
	}
	x.formatLogAndWriteChan(formatValue(txt, a...), InfoLevel)
}

func (x *XLog) Warn(txt string, a ...interface{}) {
	if x.level > WarnLevel {
		return
	}
	x.formatLogAndWriteChan(formatValue(txt, a...), WarnLevel)
}

func (x *XLog) Error(txt string, a ...interface{}) {
	if x.level > ErrorLevel {
		return
	}
	x.formatLogAndWriteChan(formatValue(txt, a...), ErrorLevel)
}

func (x *XLog) Fatal(txt string, a ...interface{}) {
	if x.level > FatalLevel {
		return
	}
	x.formatLogAndWriteChan(formatValue(txt, a...), FatalLevel)
}

func (x *XLog) formatLogAndWriteChan(body *string, level int) {
	var buffer bytes.Buffer

	buffer.WriteString("[")
	buffer.WriteString(time.Now().Format("2006-01-02 15-04-05.0000"))
	buffer.WriteString("] ")

	buffer.WriteString("[")
	buffer.WriteString(x.hostname)
	buffer.WriteString("] ")

	buffer.WriteString("[")
	buffer.WriteString(x.service)
	buffer.WriteString("] ")

	f, filename, lineno := getRuntimeInfo(x.skip)
	buffer.WriteString("[")
	buffer.WriteString(fmt.Sprintf("%s:%d%s", path.Base(filename), lineno, f))
	buffer.WriteString("] ")

	buffer.WriteString(*body)
	buffer.WriteString("\n")
	x.logTextChan <- colors[level](buffer.String())
}
