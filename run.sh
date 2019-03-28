#!/bin/sh -eu
sadm_args=${sadm_args:-''}
cfg=./etc/devel.json
go run ./bin/sadm/main.go -config $cfg $sadm_args $@
