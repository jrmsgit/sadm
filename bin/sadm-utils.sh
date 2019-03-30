#!/bin/sh -eu
cfg=./etc/devel.json
go build -o ./bin/sadm-utils.bin -i ./bin/sadm-utils
./bin/sadm-utils.bin -config $cfg $@
