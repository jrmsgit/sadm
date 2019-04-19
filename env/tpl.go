// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"path/filepath"
	"strings"

	//~ "github.com/jrmsdev/sadm/internal/log"
)

func (e *Env) TplFile(relname string) string {
	relname = filepath.Clean(relname)
	return filepath.Join(e.cfg.CfgDir, "template", relname)
}

type TplData struct {
	db map[string]string
}

func (e *Env) TplData() *TplData {
	return &TplData{e.db}
}

func (d *TplData) Get(opt string) string {
	return strings.TrimSpace(d.db[opt])
}
