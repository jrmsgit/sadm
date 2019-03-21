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
	l.Printf = printf
	l.Print = print
	if level == "debug" {
		l.Debug = debug
	} else if level == "quiet" {
		l.Debug = quietf
		l.Printf = quietf
		l.Print = quiet
	}
	return nil
}

func quietf(format string, args ...interface{}) {
}

func quiet(args ...interface{}) {
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "E: "+format+"\n", args...)
}

func perror(err error) {
	errorf("%s", err)
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

func Printf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func Print(args ...interface{}) {
	l.Print(args...)
}
