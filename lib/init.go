// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package lib

import (
	"github.com/jrmsdev/gojc/fs"
	"github.com/jrmsdev/sadm/internal/log"
)

var prefix string
var zipfile string

func Init() error {
	if prefix == "" {
		prefix = "./lib"
	}
	log.Debug("prefix %s", prefix)
	log.Debug("zipfile %d", len(zipfile))
	return fs.Init(prefix, zipfile)
}
