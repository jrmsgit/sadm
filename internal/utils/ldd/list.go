// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ldd

import (
	"github.com/jrmsdev/sadm/internal/log"
)

func List(cmd string) ([]string, error) {
	log.Debug("list %s", cmd)
	l := make([]string, 0)
	return l, nil
}
