// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"os"

	"github.com/jrmsdev/sadm/internal/cfg"
)

var prefix = "/usr/local"

func init() {
	p := os.Getenv("SADM_PREFIX")
	if p != "" {
		prefix = p
	}
}

func NewConfig(src string) *cfg.Cfg {
	x, err := cfg.New(NewReadCloser(src))
	if err != nil {
		panic(err)
	}
	x.CfgDir = prefix + "/etc/sadm"
	return x
}
