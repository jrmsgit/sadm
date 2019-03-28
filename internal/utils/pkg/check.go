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
	log.Debug("which %s: %s", filename, info.Pkg)
	err = check(m, info, info.Pkg)
	if err != nil {
		return err
	}
	log.Debug("%s depends on %d packages", info.Pkg, len(info.Deps))
	log.Debug("%s deps %v", info.Pkg, info.Deps)
	return nil
}

func check(m Manager, info *Info, pkgname string) error {
	log.Debug("check pkg %s", pkgname)
	err := m.Depends(info, pkgname)
	if err != nil {
		return err
	}
	for _, dep := range info.Deps {
		r := &Info{}
		err = check(m, r, dep.Pkg)
		if err != nil {
			if err == depDone {
				continue
			} else {
				return err
			}
		}
		if info.Pkg != "" {
			for _, d := range r.Deps {
				log.Debug("dep append %s: %s", info.Pkg, d.Pkg)
				info.Deps = append(info.Deps, d)
			}
		}
	}
	return nil
}
