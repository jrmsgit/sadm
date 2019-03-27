// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func Check(opt *args.Args, filename string) error {
	log.Debug("check %s", filename)
	if m, err := newManager(opt); err != nil {
		return err
	} else {
		if pkgname, err := m.Which(filename); err != nil {
			return err
		} else {
			log.Debug("%s provided by %s", filename, pkgname)
		}
	}
	return nil
}
