#!/bin/sh -eu
rm -f internal/cfg/build.go
./run.sh -log debug $@
