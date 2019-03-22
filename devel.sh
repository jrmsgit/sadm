#!/bin/sh -eu
go run ./bin/sadm/main.go -log debug -config ./etc/devel/sadm.json $@
