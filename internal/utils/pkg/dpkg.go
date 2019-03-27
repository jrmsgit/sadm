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
