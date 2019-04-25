#!/bin/sh -eu
cfg=./etc/devel.json
cd internal/log && {
	cat debug.go.enable >debug.go
}
cd - >/dev/null
cd lib && {
	test -s zip.go || go run ./_zip/main.go --prefix ${PWD}
}
cd - >/dev/null
go build -i -o ./bin/sadm-utils.bin ./bin/sadm-utils
./bin/sadm-utils.bin -config ${cfg} $@
