// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"os"

	"github.com/jrmsdev/sadm/internal/log"
)

func Chmod(dst string, f os.FileMode) error {
	s, err := os.Stat(dst)
	if err != nil {
		log.Debug("%s", err)
		return err
	}
	if s.Mode().Perm() != f.Perm() {
		err = os.Chmod(dst, f.Perm())
		if err != nil {
			log.Debug("%s", err)
			return err
		}
		log.Printf("chmod %o %s", f.Perm(), dst)
	}
	return nil
}
