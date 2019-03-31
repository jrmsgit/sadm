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
rm -f internal/cfg/build.go
go test $verbose $coverage ./...
if test "X${coverage}" != 'X'; then
  go tool cover -html coverage.out -o coverage.html
fi
