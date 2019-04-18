#!/bin/sh -eu
cfg=./etc/devel.json
cd internal/log && {
	cat debug.go.enable >debug.go
}
cd - >/dev/null
go build -i -o ./bin/sadm-utils.bin ./bin/sadm-utils
./bin/sadm-utils.bin -config ${cfg} $@
