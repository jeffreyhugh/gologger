package gologger

import (
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestInit(t *testing.T) {
	Init(logrus.InfoLevel)
}

func TestDebug(t *testing.T) {
	// without logrus.Fields
	Debug("This is a Debug message", nil)
	// with logrus.Fields
	Debug("This is a Debug message", logrus.Fields{"log_message": "Here's a Debug message"})
}

func TestInfo(t *testing.T) {
	Info("This is an Info message", nil)
	Info("This is an Info message", logrus.Fields{"log_message": "Here's an Info message"})
}

func TestWarn(t *testing.T) {
	err := errors.New("This is an error from Warn")
	// test with no error or fields
	Warn("This is a Warn message", nil, nil)
	// error and no fields
	Warn("This is a Warn message", err, nil)
	// error and fields
	Warn("This is a Warn message", err, logrus.Fields{"log_message": "Here's a Warn message"})

}

func TestWarning(t *testing.T) {
	err := errors.New("This is an error from Warning")
	Warning("This is a Warn message", nil, nil)
	Warning("This is a Warn message", err, nil)
	Warning("This is a Warn message", err, logrus.Fields{"log_message": "Here's a Warning message"})
}

func TestError(t *testing.T) {
	err := errors.New("This is an error from Error")
	Error("This is an Error message", nil, nil)
	Error("This is an Error message", err, nil)
	Error("This is an Error message", err, logrus.Fields{"log_message": "Here's an Error message"})
}

// Unfortunately, there doesn't seem to be a good way to test Fatal(), as it exits and if tested directly will
// make the whole test fail. There is a solution out there: https://stackoverflow.com/questions/26225513/how-to-test-os-exit-scenarios-in-go
// However, these passed tests won't count towards the percentage of passed tests.
