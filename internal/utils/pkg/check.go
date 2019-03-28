// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func Check(opt *args.Args, filename string) error {
	log.Debug("check %s", filename)
	var (
		m       Manager
		err     error
		pkgname string
		deps    []string
	)
	m, err = newManager(opt)
	if err != nil {
		return err
	}
	pkgname, err = m.Which(filename)
	if err != nil {
		return err
	}
	log.Debug("%s provided by %s", filename, pkgname)
	deps, err = m.Depends(pkgname)
	if err != nil {
		return err
	}
	log.Debug("%s depends on %v", pkgname, deps)
	return nil
}
