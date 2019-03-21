// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package args

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
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
	log.Debug("new")
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

func (a *Args) load(prefix string, fh io.ReadCloser) {
	defer fh.Close()
	src := make(map[string]string)
	if blob, err := ioutil.ReadAll(fh); err != nil {
		log.Warn(err)
	} else {
		if err := json.Unmarshal(blob, &src); err != nil {
			log.Warn(err)
		} else {
			for opt, val := range src {
				a.db[prefix+"."+opt] = val
			}
			//~ log.Debug("%s loaded %#v", prefix, a.db)
		}
	}
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
			if fh, err := os.Open(fn); err != nil {
				log.Warn(err)
			} else {
				a.load(s, fh)
			}
		}
	}
}

func (a *Args) Get(opt string) string {
	return a.db[opt]
}
