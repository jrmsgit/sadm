#!/bin/sh -eu
gen=../../gojc/fs/_zipfile/gen.go
if test -s ../vendor/github.com/jrmsdev/gojc/fs/_zipfile/gen.go; then
	gen=../vendor/github.com/jrmsdev/gojc/fs/_zipfile/gen.go
fi
echo "run ${gen}"
go run ${gen} $@ '*/*.json' '*/*/*.json'
