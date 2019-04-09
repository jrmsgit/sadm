// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"os"

	"github.com/jrmsdev/sadm/internal/log"
)

func Exists(filename string) bool {
	s, err := os.Stat(filename)
	if err == nil {
		if s.IsDir() {
			log.Debug("%s dir found", filename)
		} else {
			log.Debug("%s file found", filename)
		}
		return true
	}
	return false
}
