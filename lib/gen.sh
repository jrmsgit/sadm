#!/bin/sh -eu
go run ../../gojc/fs/_zipfile/gen.go $@ '*/*.json' '*/*/*.json'
