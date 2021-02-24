# gologger

## Purpose

`gologger` provides a simple wrapper around the `logrus` logging package in an intuitive manner. Using this 
package in conjunction with `logrus` should slim the codebase down and be a little more intuitive.

## Installation
`go get -d -v github.com/qbxt/gologger`

## Example

```go
package main

import (
    "errors"

    "github.com/qbxt/gologger"
    "github.com/sirupsen/logrus"
)

func main() {
    gologger.Init(logrus.InfoLevel)

    gologger.Debug("This is a Debug message", nil)
    gologger.Debug("This is a Debug message", logrus.Fields{"log_message": "Here's a Debug message"})
    gologger.Info("This is an Info message", logrus.Fields{"log_message": "Here's an Info message"})

    err := errors.New("This is a generic error!")
    gologger.Warn("This is a Warn message", err, nil)
    gologger.Error("This is an Error message", err, nil)
    gologger.Fatal("This is a Fatal message", err, logrus.Fields{"log_message": "Something is very very wrong"})
}
```