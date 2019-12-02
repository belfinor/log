package log

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2017-12-02

import (
	"testing"
)

func TestLoggerSetLevel(t *testing.T) {

	defLog, _ = New(&Config{}, true)

	if GetLevel() != "none" {
		t.Fatal("invalid default log level")
	}

	SetLevel("info")

	if GetLevel() != "info" {
		t.Fatal("expected log level 'info'")
	}

	Errorf("failed %d", 1)
	Finishf("finish %s", "ok")

	defLog.Close()

	defLog = nil

	Error("error1", "error2")
	Warn("warn")
	Info("info")
	Debug("debug")
	Trace("trace", 1, "trace", 2)
	Infof("Hello, %s", "Mike")
	Errorf("Hello, %s", "Mike")
	Warnf("Hello, %s", "Mike")
	Debugf("Hello, %s", "Mike")
	Tracef("Hello, %s", "Mike")

	defLog.Close()
}
