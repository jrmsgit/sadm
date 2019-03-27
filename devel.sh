#!/bin/sh -eu
rm -f internal/cfg/build.go
go run ./bin/sadm/main.go -log debug -config ./etc/devel/sadm.json $@
