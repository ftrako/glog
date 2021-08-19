package glog

import (
	"log"
	"testing"
	"time"
)

func BenchmarkTrace(b *testing.B) {
	SetFlag(log.Lshortfile | log.Lmicroseconds)
	// EnableFuncName(false)
	// EnablePrefix(true)
	// EnableColor(false)
	now := time.Now()
	for i := 0; i < b.N; i++ {
		Debug("abc")
	}
	Debug("size:", b.N)
	Debug("cost:", time.Since(now))
}

func TestInfo(t *testing.T) {
	// EnableColor(true)
	Debug("abc_debug")
	Info("abc_info")
	Warn("abc_warn")
	Error("abc_err")
	// Panic("abc_painc")
}

func TestInfoDepth(t *testing.T) {
	a1()
}

func a1() {
	b1()
}

func b1() {
	Info("abc")
}
