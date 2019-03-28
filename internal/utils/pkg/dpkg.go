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
	cmd  string
	arch string
}

func dpkgNew(opt *args.Args) *dpkgManager {
	cmd := opt.Get("os.pkg.exec")
	arch := opt.Get("arch")
	log.Debug("dpkg new %s (arch:%s)", cmd, arch)
	return &dpkgManager{cmd, arch}
}

func (m *dpkgManager) Which(info *Info, filename string) error {
	if out, err := utils.Exec(m.cmd, "-S", filename); err != nil {
		return err
	} else {
		info.Pkg = strings.Split(strings.TrimSpace(string(out)), ":")[0]
	}
	return nil
}

func (m *dpkgManager) Depends(info *Info) error {
	if info.Deps == nil {
		info.Deps = make([]*Info, 0)
	}
	out, err := utils.Exec(m.cmd+"-query", "-W", "-f ${Depends}", info.Pkg+":"+m.arch)
	if err != nil {
		return err
	}
	log.Debug("%s", out)
	for _, line := range strings.Split(string(out), ",") {
		n := strings.Split(strings.TrimSpace(line), " ")[0]
		log.Debug("%s", n)
		i := &Info{}
		i.Pkg = n
		info.Deps = append(info.Deps, i)
	}
	return nil
}
