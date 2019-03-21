// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"fmt"
	"os"
	"runtime"
)

type logger struct {
	Debug  func(string, ...interface{})
	Errorf func(string, ...interface{})
	Error  func(error)
	Warnf  func(string, ...interface{})
	Warn   func(error)
	Printf func(string, ...interface{})
	Print  func(...interface{})
}

var (
	l *logger
)

func Init(level string) error {
	l = new(logger)
	l.Debug = quietf
	l.Errorf = errorf
	l.Error = perror
	l.Warnf = quietf
	l.Warn = quieterr
	l.Printf = printf
	l.Print = print
	if level == "debug" {
		l.Debug = debug
		l.Warnf = warnf
		l.Warn = warn
	} else if level == "warn" {
		l.Warnf = warnf
		l.Warn = warn
	} else if level == "quiet" {
		l.Printf = quietf
		l.Print = quiet
	}
	return nil
}

func quietf(format string, args ...interface{}) {
}

func quiet(args ...interface{}) {
}

func quieterr(err error) {
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "E: "+format+"\n", args...)
}

func perror(err error) {
	errorf("%s", err)
}

func warnf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "W: "+format+"\n", args...)
}

func warn(err error) {
	warnf("%s", err)
}

func printf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func print(args ...interface{}) {
	fmt.Print(args...)
	fmt.Print("\n")
}

func debug(format string, args ...interface{}) {
	tag := "D: "
	_, fn, ln, ok := runtime.Caller(2)
	if ok {
		tag = fmt.Sprintf("%s:%d ", fn, ln)
	}
	fmt.Fprintf(os.Stderr, tag+format+"\n", args...)
}

func Debug(format string, args ...interface{}) {
	l.Debug(format, args...)
}

func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func Error(err error) {
	l.Error(err)
}

func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func Warn(err error) {
	l.Warn(err)
}

func Printf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func Print(args ...interface{}) {
	l.Print(args...)
}
