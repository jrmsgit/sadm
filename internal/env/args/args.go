// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package args

import (
	"path/filepath"
	"runtime"

	"github.com/jrmsdev/sadm/internal/cfg"
	"github.com/jrmsdev/sadm/internal/log"
)

type Args struct {
	cfg *cfg.Cfg
	db  map[string]string
}

func New(config *cfg.Cfg, src map[string]string) *Args {
	a := &Args{config, src}
	a.setRuntime()
	a.setOS()
	a.loadService()
	//~ log.Debug("new %#v", a)
	return a
}

func (a *Args) setRuntime() {
	a.db["os"] = runtime.GOOS
	a.db["arch"] = runtime.GOARCH
}

func (a *Args) setOS() {
	a.db["pkgman"] = "dpkg"
}

func (a *Args) loadService() {
	s := a.db["service"]
	if s != "" {
		files := []string{
			filepath.Join(a.cfg.LibDir, "service", s, "config.json"),
			filepath.Join(a.cfg.LibDir, "service", s, "config-"+a.db["os"]+".json"),
		}
		for _, fn := range files {
			log.Debug("service load %s", fn)
		}
	}
}

func (a *Args) Get(opt string) string {
	return a.db[opt]
}
