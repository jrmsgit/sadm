#!/bin/sh -eu
cfg=./etc/devel.json
go build -i -o ./bin/sadm-utils.bin ./bin/sadm-utils
./bin/sadm-utils.bin -config $cfg $@
