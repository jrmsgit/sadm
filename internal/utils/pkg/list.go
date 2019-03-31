// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	"strings"

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
	// list deps
	pkgDeps := strings.Split(strings.TrimSpace(opt.Get("service.pkg.deps")), " ")
	//~ log.Debug("pkg.deps %s: %v", info.Pkg, pkgDeps)
	err = getDeps(m, info, info.Pkg, pkgDeps...)
	if err != nil {
		return err
	}
	log.Debug("%s requires %d packages", info.Pkg, len(info.Deps))
	// list files
	err = getFiles(m, info, info.Pkg)
	if err != nil {
		return err
	}
	log.Debug("%s requires %d files", info.Pkg, len(info.Files))
	return nil
}

func getDeps(m Manager, info *Info, pkgname string, pkgdeps ...string) error {
	//~ log.Debug("get deps: %s", pkgname)
	err := m.Depends(info, pkgname)
	if err != nil {
		return err
	}
	for _, dep := range info.Deps {
		if err := lsDeps(m, info, dep.Pkg); err != nil {
			return err
		}
	}
	for _, n := range pkgdeps {
		n = strings.TrimSpace(n)
		if n == "" {
			continue
		}
		if err := lsDeps(m, info, n); err != nil {
			return err
		} else {
			d := new(Info)
			d.Pkg = n
			info.Deps = append(info.Deps, d)
		}
	}
	return nil
}

func lsDeps(m Manager, info *Info, pkgname string) error {
	//~ log.Debug("lsdeps: %s (%s)", pkgname, info.Pkg)
	r := &Info{}
	err := getDeps(m, r, pkgname)
	if err != nil {
		if err == depDone {
			return nil
		} else {
			return err
		}
	}
	for _, d := range r.Deps {
		//~ log.Debug("%s <- %s", pkgname, d.Pkg)
		info.Deps = append(info.Deps, d)
	}
	return nil
}

func getFiles(m Manager, info *Info, pkgname string) error {
	//~ log.Debug("get files: %s", pkgname)
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
