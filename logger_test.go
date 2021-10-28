package log

import (
	"testing"
	"time"
)

func TestInitLogger(t *testing.T) {
	//InitLogger(NewFileLogger())
	for i := 0; i < 1000; i++ {
		Debug("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
		Info("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
		Warn("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
		Error("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
		Fatal("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
		time.Sleep(1 * time.Second)
	}
}
