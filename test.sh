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
test -f ./lib/zipfs.go || (cd ./lib && ./gen.sh --prefix ${PWD})
go test $verbose $coverage $tpath
if test "X${coverage}" != 'X'; then
  go tool cover -html coverage.out -o coverage.html
fi
