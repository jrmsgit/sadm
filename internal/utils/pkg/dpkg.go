// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	"strings"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils"
)

type dpkgManager struct {
	cmd     string
	arch    string
	exclude map[string]bool
	depdone map[string]bool
	deps    map[string]bool
	prune   map[string]bool
	files   map[string]bool
}

func dpkgNew(opt *args.Args) *dpkgManager {
	cmd := opt.Get("os.pkg.exec")
	arch := opt.Get("arch")
	log.Debug("dpkg new %s (arch:%s)", cmd, arch)
	return &dpkgManager{
		cmd,
		arch,
		pkgExclude(opt),
		make(map[string]bool),
		make(map[string]bool),
		filesPrune(opt),
		make(map[string]bool),
	}
}

// load pkg exclude config

func pkgExclude(opt *args.Args) map[string]bool {
	excl := make(map[string]bool)
	for _, n := range strings.Split(opt.Get("os.pkg.exclude"), " ") {
		n = strings.TrimSpace(n)
		if n != "" {
			excl[n] = true
		}
	}
	return excl
}

// load files prune config

func filesPrune(opt *args.Args) map[string]bool {
	prune := make(map[string]bool)
	for _, n := range strings.Split(opt.Get("os.pkg.prune"), " ") {
		n = strings.TrimSpace(n)
		if n != "" {
			prune[n] = true
		}
	}
	return prune
}

// which package provides filename?

func (m *dpkgManager) Which(info *Info, filename string) error {
	if out, err := utils.Exec(m.cmd, "-S", filename); err != nil {
		return err
	} else {
		info.Pkg = m.fullname(strings.Split(strings.TrimSpace(string(out)), ":")[0])
	}
	return nil
}

// find package dependencies

func (m *dpkgManager) Depends(info *Info, pkgname string) error {
	log.Debug("find deps: %s", pkgname)
	if m.depdone[pkgname] {
		//~ log.Debug("deps done: %s", pkgname)
		return depDone
	}
	if info.Deps == nil {
		info.Deps = make([]*Info, 0)
	}
	out, err := utils.Exec(m.cmd+"-query", "--no-pager", "-W", "-f", "${Depends}", pkgname)
	if err != nil {
		return err
	}
	//~ log.Debug("deps %s: %s", pkgname, out)
	for _, line := range strings.Split(string(out), ",") {
		n := strings.Split(strings.TrimSpace(line), " ")[0]
		fulln := m.fullname(n)
		if n == "" || fulln == "" {
			continue
		}
		_, excl := m.exclude[n]
		if excl {
			log.Warnf("pkg exclude %s (%s)", fulln, n)
			continue
		} else {
			added := m.deps[fulln]
			if !added {
				i := &Info{}
				i.Pkg = fulln
				info.Deps = append(info.Deps, i)
				m.deps[fulln] = true
			}
		}
	}
	m.depdone[pkgname] = true
	return nil
}

// find package fullname (ie: name:arch)

func (m *dpkgManager) fullname(pkgname string) string {
	if pkgname == "" {
		return ""
	}
	out, err := utils.Exec(m.cmd+"-query", "-W", "-f", "${binary:Package}\n", pkgname)
	if err != nil {
		log.Warn(err)
		return pkgname
	}
	outs := string(out)
	if len(strings.Split(outs, "\n")) > 2 {
		return pkgname + ":" + m.arch
	}
	return strings.TrimSpace(outs)
}

// list files provided by package

func (m *dpkgManager) List(info *Info, pkgname string) error {
	log.Debug("list %s", pkgname)
	out, err := utils.Exec(m.cmd, "-L", pkgname)
	if err != nil {
		return err
	}
	count := 0
	pruneCount := 0
	for _, fn := range strings.Split(string(out), "\n") {
		fn = strings.TrimSpace(fn)
		if fn == "" {
			continue
		}
		if !m.files[fn] {
			if m.pruneFile(fn) {
				log.Warnf("pkg prune %s", fn)
				pruneCount += 1
			} else {
				log.Debug("file append: %s", fn)
				info.Files = append(info.Files, fn)
				m.files[fn] = true
				count += 1
			}
		}
	}
	log.Debug("append count %s: %d", pkgname, count)
	log.Debug("prune count %s: %d", pkgname, pruneCount)
	return nil
}

func (m *dpkgManager) pruneFile(name string) bool {
	for prune := range m.prune {
		if strings.HasPrefix(name, prune) {
			return true
		}
	}
	return false
}
