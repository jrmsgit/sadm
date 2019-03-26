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
	OS string
	Service string
}

func New(config *cfg.Cfg, src map[string]string) (*Args, error) {
	log.Debug("new")
	a := new(Args)
	a.cfg = config
	a.db = src
	a.OS = runtime.GOOS
	a.setRuntime()
	if err := a.loadOS(); err != nil {
		return nil, err
	}
	a.Service = src["service"]
	if err := a.loadService(); err != nil {
		return nil, err
	}
	//~ log.Debug("new %#v", a)
	return a, nil
}

func (a *Args) setRuntime() {
	a.db["os"] = a.OS
	a.db["arch"] = runtime.GOARCH
}

func (a *Args) load(prefix string, fh io.ReadCloser) error {
	defer fh.Close()
	src := make(map[string]string)
	if blob, err := ioutil.ReadAll(fh); err != nil {
		log.Debug("%s", err)
		return err
	} else {
		if err := json.Unmarshal(blob, &src); err != nil {
			log.Debug("%s", err)
			return err
		} else {
			if prefix != "" {
				prefix = prefix + "."
			}
			for opt, val := range src {
				a.db[prefix+opt] = val
			}
			//~ log.Debug("%s loaded %#v", prefix, a.db)
		}
	}
	return nil
}

func (a *Args) loadOS() error {
	fn := filepath.Join(a.cfg.CfgDir, "os", "config.json")
	if fh, err := os.Open(fn); err != nil {
		log.Debug("%s", err)
		return err
	} else {
		log.Debug("load %s", fn)
		if err := a.load("os", fh); err != nil {
			return err
		}
	}
	n := a.OS
	if n != "" {
		fn := filepath.Join(a.cfg.CfgDir, "os", n, "config.json")
		if fh, err := os.Open(fn); err != nil {
			log.Debug("%s", err)
			return err
		} else {
			log.Debug("load %s", fn)
			return a.load("os", fh)
		}
	}
	return nil
}

func (a *Args) loadService() error {
	fn := filepath.Join(a.cfg.CfgDir, "service", "config.json")
	if fh, err := os.Open(fn); err != nil {
		log.Debug("%s", err)
		return err
	} else {
		log.Debug("load %s", fn)
		if err := a.load("service", fh); err != nil {
			return err
		}
	}
	s := a.Service
	if s != "" {
		files := []string{
			//~ filepath.Join(a.cfg.CfgDir, "service", s, "config.json"),
			filepath.Join(a.cfg.CfgDir, "service", s, a.db["os"]+".json"),
		}
		for _, fn := range files {
			if fh, err := os.Open(fn); err != nil {
				log.Debug("%s", err)
				return err
			} else {
				log.Debug("load %s", fn)
				if err := a.load("service", fh); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (a *Args) Get(opt string) string {
	return a.db[opt]
}
