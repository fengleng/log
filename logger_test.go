package log

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	InitLogger(nil)
	Debug("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	Info("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	Warn("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	Error("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
	Fatal("fasfaf%s", "gsgggggggggggggggggggggggggggggsgsgsgsg")
}
