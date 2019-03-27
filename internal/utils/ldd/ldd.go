// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ldd

import (
	"github.com/jrmsdev/sadm/internal/log"
)

type Info struct {
	Filename string
	Files    []string
}

func newInfo(filename string) *Info {
	log.Debug("new info %s", filename)
	return &Info{
		filename,
		make([]string, 0),
	}
}
