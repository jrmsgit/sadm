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
		m   Manager
		err error
	)
	m, err = newManager(opt)
	if err != nil {
		return err
	}
	info := &Info{}
	err = m.Which(info, filename)
	if err != nil {
		return err
	}
	log.Debug("%s provided by %s", filename, info.Pkg)
	return check(m, info, info.Pkg)
}

func check(m Manager, info *Info, pkgname string) error {
	err := m.Depends(info, pkgname)
	if err != nil {
		return err
	}
	for _, dep := range info.Deps {
		err = check(m, info, dep.Pkg)
		if err != nil {
			return err
		}
	}
	log.Debug("%s depends on %d packages", info.Pkg, len(info.Deps))
	return nil
}
