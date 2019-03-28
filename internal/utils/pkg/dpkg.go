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
	cmd string
}

func dpkgNew(opt *args.Args) *dpkgManager {
	cmd := opt.Get("os.pkg.exec")
	log.Debug("dpkg new %s", cmd)
	return &dpkgManager{cmd}
}

func (m *dpkgManager) Which(filename string) (string, error) {
	n := ""
	if out, err := utils.Exec(m.cmd, "-S", filename); err != nil {
		return "", err
	} else {
		return strings.Split(strings.TrimSpace(string(out)), ":")[0], nil
	}
	return n, nil
}

func (m *dpkgManager) Depends(name string) ([]string, error) {
	l := make([]string, 0)
	out, err := utils.Exec(m.cmd+"-query", "-W", "-f ${Depends}", name)
	if err != nil {
		return nil, err
	}
	log.Debug("%s", out)
	for _, line := range strings.Split(string(out), ",") {
		n := strings.Split(strings.TrimSpace(line), " ")[0]
		log.Debug("%s", n)
	}
	return l, nil
}
