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
}

func dpkgNew(opt *args.Args) *dpkgManager {
	cmd := opt.Get("os.pkg.exec")
	arch := opt.Get("arch")
	log.Debug("dpkg new %s (arch:%s)", cmd, arch)
	return &dpkgManager{cmd, arch, pkgExclude(opt),
		make(map[string]bool), make(map[string]bool)}
}

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

func (m *dpkgManager) Which(info *Info, filename string) error {
	if out, err := utils.Exec(m.cmd, "-S", filename); err != nil {
		return err
	} else {
		info.Pkg = m.fullname(strings.Split(strings.TrimSpace(string(out)), ":")[0])
	}
	return nil
}

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
