#!/bin/sh -eu
verbose=''
if test "${1:-'none'}" = '-v'; then
  shift
  verbose='-v'
fi
coverage=''
if test "${1:-'none'}" = '--coverage'; then
  shift
  coverage='-coverprofile coverage.out'
fi
tpath=${1:-'./...'}
export SADM_PREFIX=${PWD}
rm -f internal/cfg/build.go
rm -f internal/log/debug.go
test -f ./lib/zip.go || (cd lib && go run ./_zip/main.go)
go test $verbose $coverage $tpath
if test "X${coverage}" != 'X'; then
  go tool cover -html coverage.out -o coverage.html
fi
