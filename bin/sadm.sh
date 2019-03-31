#!/bin/sh -eu
cfg=./etc/devel.json
go build -i -o ./bin/sadm.bin ./bin/sadm
./bin/sadm.bin -config $cfg $@
