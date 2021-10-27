package log

import "testing"

func TestColor(t *testing.T) {
	c := len(colors)
	for i := 0; i < c; i++ {
		t.Log(colors[i]("ccccccccc"))
	}
}
