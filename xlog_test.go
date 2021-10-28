package log

import (
	"testing"
)

func TestNewXLog(t *testing.T) {
	//time.Sleep(10 * time.Second)
}

func TestNewConsoleLogger(t *testing.T) {
	log := NewConsoleLogger()
	//log := xlog.NewXLog(nil)
	log.Debug("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Info("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Warn("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Error("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Fatal("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
}
