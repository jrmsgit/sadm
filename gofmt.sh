#!/bin/sh -eu
gofmt -w -l -s .
gofmt -w -l -s internal/log/debug.go.enable
