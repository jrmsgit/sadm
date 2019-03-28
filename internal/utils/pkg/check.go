// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func Check(opt *args.Args, filename string) (*Info, error) {
	log.Debug("check %s", filename)
	var (
		m   Manager
		err error
	)
	m, err = newManager(opt)
	if err != nil {
		return nil, err
	}
	info := &Info{}
	err = m.Which(info, filename)
	if err != nil {
		return nil, err
	}
	log.Debug("which %s: %s", filename, info.Pkg)
	err = getDeps(m, info, info.Pkg)
	if err != nil {
		return nil, err
	}
	log.Debug("%s requires %d packages", info.Pkg, len(info.Deps))
	//~ log.Debug("%s deps %v", info.Pkg, info.Deps)
	//~ info.Files = make([]string, 0)
	err = getFiles(m, info, info.Pkg)
	if err != nil {
		return nil, err
	}
	log.Debug("%s requires %d files", info.Pkg, len(info.Files))
	//~ log.Debug("%s files %v", info.Pkg, info.Files)
	//~ for _, fn := range info.Files {
	//~ log.Print(fn)
	//~ }
	return info, nil
}

func getDeps(m Manager, info *Info, pkgname string) error {
	//~ log.Debug("get deps %s", pkgname)
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
				//~ log.Debug("dep append %s: %s", info.Pkg, d.Pkg)
				info.Deps = append(info.Deps, d)
			}
		}
	}
	return nil
}

func getFiles(m Manager, info *Info, pkgname string) error {
	//~ log.Debug("get files %s", pkgname)
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
