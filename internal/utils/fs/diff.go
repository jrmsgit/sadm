// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"github.com/jrmsdev/sadm/internal/log"
)

func diff(src, dst *Info) bool {
	log.Debug("diff %s %s", src, dst)
	fail := false
	isdir := src.IsDir()
	if src.IsDir() && !dst.IsDir() {
		log.Warnf("diff type\n      is dir %s\n     not dir %s", src, dst)
		fail = true
	}
	if !src.IsDir() && dst.IsDir() {
		log.Warnf("diff type\n     not dir %s\n      is dir %s", src, dst)
		fail = true
	}
	if !isdir && (src.Size() != dst.Size()) {
		log.Warnf("diff size\n     %d %s\n     %d %s", src.Size(), src, dst.Size(), dst)
		fail = true
	}
	if src.Mode().Perm() != dst.Mode().Perm() {
		log.Warnf("diff mode\n     %o %s\n     %o %s",
			src.Mode().Perm(), src, dst.Mode().Perm(), dst)
		fail = true
	}
	return fail
}
