// Package gologger package is used as a simple wrapper for the logrus
// logging package. This simple layer of abstraction requires less
// code and can keep code much cleaner in the end.
package gologger

// only import needed is logrus
import (
	"github.com/sirupsen/logrus"
)

// Init globally sets the logger level to the provided logrus.Level - which can be any of:
//  * logrus.PanicLevel
//  * logrus.FatalLevel
//  * logrus.ErrorLevel
//  * logrus.WarnLevel
//  * logrus.InfoLevel
//  * logrus.DebugLevel
func Init(l logrus.Level) {
	logrus.SetLevel(l)
}

// Debug writes a message to the log of Debug level status. Requires:
//  * message (string): helpful, user friendly error text
//  * fields (logrus.Fields): Can be nil. If set, provide these logrus.Fields to the entry
func Debug(message string, fields logrus.Fields) {
	if fields != nil {
		logrus.WithFields(fields).Debug(message)
	} else {
		logrus.Debug(message)
	}
}

// Info writes a message to the log of Info level status. Requires:
//  * message (string): helpful, user friendly error text
//  * fields (logrus.Fields): Can be nil. If set, provide these logrus.Fields to the entry
func Info(message string, fields logrus.Fields) {
	if fields != nil {
		logrus.WithFields(fields).Info(message)
	} else {
		logrus.Info(message)
	}
}

// Warn writes a message to the log of Warn level status. Requires:
//  * message (string): helpful, user friendly error text
//  * err (error): An error obtained from a failed call to a previous method or function
//  * fields (logrus.Fields): Can be nil. If set, provide these logrus.Fields to the entry
func Warn(message string, err error, fields logrus.Fields) {
	if fields != nil {
		fields["error"] = err
		logrus.WithFields(fields).Warn(message)
	} else if err != nil {
		fields := logrus.Fields{"error": err}
		logrus.WithFields(fields).Warn(message)
	} else {
		logrus.Warn(message)
	}
}

// Warning is an alias for Warn. Will call Warn() with provided options
func Warning(message string, err error, fields logrus.Fields) {
	Warn(message, err, fields)
}

// Error writes a message to the log of Error level status. Requires:
//  * message (string): helpful, user friendly error text
//  * err (error): An error obtained from a failed call to a previous method or function
//  * fields (logrus.Fields): Can be nil. If set, provide these logrus.Fields to the entry
func Error(message string, err error, fields logrus.Fields) {
	if fields != nil {
		fields["error"] = err
		logrus.WithFields(fields).Error(message)
	} else if err != nil {
		fields := logrus.Fields{"error": err}
		logrus.WithFields(fields).Error(message)
	} else {
		logrus.Error(message)
	}
}

// Fatal writes a message to the log of Fatal level status. Requires:
//  * message (string): helpful, user friendly error text
//  * err (error): An error obtained from a failed call to a previous method or function
//  * fields (logrus.Fields): Can be nil. If set, provide these logrus.Fields to the entry
// Note: Calling a Fatal() error will exit execution of the current program. Goroutines will not
// execute on deferral. Only call Fatal() if you are sure that the program should exit as well.
func Fatal(message string, err error, fields logrus.Fields) {
	if fields != nil {
		fields["error"] = err
		logrus.WithFields(fields).Fatal(message)
	} else if err != nil {
		fields := logrus.Fields{"error": err}
		logrus.WithFields(fields).Fatal(message)
	} else {
		logrus.Fatal(message)
	}
}
