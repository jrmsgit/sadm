// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"github.com/jrmsdev/sadm/internal/cfg"
)

func NewConfig(src string) *cfg.Cfg {
	x, err := cfg.New(NewReadCloser(src))
	if err != nil {
		panic(err)
	}
	return x
}
