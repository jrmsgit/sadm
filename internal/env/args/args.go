// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package args

import (
	"github.com/jrmsdev/sadm/internal/log"
)

type Args struct {
	src map[string]string
}

func New(src map[string]string) *Args {
	log.Debug("new")
	return &Args{src}
}

func (a *Args) Get(opt string) string {
	return a.src[opt]
}
