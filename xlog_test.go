package log

import (
	"testing"
)

func TestNewXLog(t *testing.T) {
	log := NewXLog(nil)
	//log := xlog.NewXLog(nil)
	log.Debug("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Info("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Warn("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Error("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Fatal("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	//time.Sleep(10 * time.Second)
}

func TestNewConsoleLogger(t *testing.T) {
	log := NewConsoleLogger(nil)
	//log := xlog.NewXLog(nil)
	log.Debug("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Info("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Warn("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Error("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	log.Fatal("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
}
