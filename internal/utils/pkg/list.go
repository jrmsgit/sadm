// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func List(opt *args.Args, info *Info) error {
	log.Debug("list %s", info.Pkg)
	var (
		m   Manager
		err error
	)
	m, err = newManager(opt)
	if err != nil {
		return err
	}
	err = getDeps(m, info, info.Pkg)
	if err != nil {
		return err
	}
	log.Debug("%s requires %d packages", info.Pkg, len(info.Deps))
	err = getFiles(m, info, info.Pkg)
	if err != nil {
		return err
	}
	log.Debug("%s requires %d files", info.Pkg, len(info.Files))
	return nil
}

func getDeps(m Manager, info *Info, pkgname string) error {
	err := m.Depends(info, pkgname)
	if err != nil {
		return err
	}
	for _, dep := range info.Deps {
		r := &Info{}
		err = getDeps(m, r, dep.Pkg)
		if err != nil {
			if err == depDone {
				continue
			} else {
				return err
			}
		}
		if info.Pkg != "" {
			for _, d := range r.Deps {
				info.Deps = append(info.Deps, d)
			}
		}
	}
	return nil
}

func getFiles(m Manager, info *Info, pkgname string) error {
	err := m.List(info, pkgname)
	if err != nil {
		return err
	}
	for _, dep := range info.Deps {
		r := &Info{}
		err = getFiles(m, r, dep.Pkg)
		if err != nil {
			return err
		}
		info.Files = append(info.Files, r.Files...)
	}
	return nil
}
