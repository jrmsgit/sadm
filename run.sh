#!/bin/sh -eu
sadm_args=${sadm_args:-''}
cfg=./etc/devel.json
GOBIN=${PWD}/build go install -i ./bin/sadm
./build/sadm -config $cfg $sadm_args $@
