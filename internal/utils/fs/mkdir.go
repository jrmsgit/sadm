// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"errors"
	"os"

	"github.com/jrmsdev/sadm/internal/log"
)

func Mkdir(dirname string) error {
	s, err := os.Stat(dirname)
	if err != nil {
		if err := os.MkdirAll(dirname, 0755); err != nil {
			log.Debug("%s", err)
			return err
		}
		log.Printf("mkdir %s", dirname)
		return nil
	}
	if !s.IsDir() {
		return errors.New(sprintf("%s exists but is not a dir", dirname))
	}
	return nil
}
