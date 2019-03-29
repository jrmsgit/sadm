// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"github.com/jrmsdev/sadm/internal/log"
)

func diff(src, dst *Info) bool {
	fail := false
	isdir := src.IsDir()
	if src.IsDir() && !dst.IsDir() {
		log.Printf("diff type is dir %s not dir %s", src, dst)
		fail = true
	}
	if !src.IsDir() && dst.IsDir() {
		log.Printf("diff type not dir %s is dir %s", src, dst)
		fail = true
	}
	if !isdir && (src.Size() != dst.Size()) {
		log.Printf("diff size %d %s %d %s", src.Size(), src, dst.Size(), dst)
		fail = true
	}
	if src.Mode().Perm() != dst.Mode().Perm() {
		log.Printf("diff mode %o %s %o %s",
			src.Mode().Perm(), src, dst.Mode().Perm(), dst)
		fail = true
	}
	if !isdir {
		ssum := checksum(src.Filename())
		dsum := checksum(dst.Filename())
		if ssum != dsum {
			log.Printf("diff checksum %s %s %s %s", ssum, src, dsum, dst)
			fail = true
		}
	}
	return fail
}
