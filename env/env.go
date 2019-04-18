// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/jrmsdev/sadm/internal/cfg"
	"github.com/jrmsdev/sadm/internal/log"
)

type Env struct {
	cfg     *cfg.Cfg
	db      map[string]string
	Name    string
	Type    string
	OS      string
	Service string
}

func New(config *cfg.Cfg, name string, src map[string]string) (*Env, error) {
	log.Debug("new %s", name)
	a := new(Env)
	a.cfg = config
	a.db = make(map[string]string)
	a.Name = name
	a.setRuntime()
	a.Type = strings.TrimSpace(src["type"])
	if err := a.init(); err != nil {
		return nil, err
	}
	if err := a.loadOS(); err != nil {
		return nil, err
	}
	a.Service = strings.TrimSpace(src["service"])
	if a.Service != "" {
		if err := a.loadService(); err != nil {
			return nil, err
		}
	}
	a.source(src)
	//~ log.Debug("new %#v", a)
	return a, nil
}

func (a *Env) source(src map[string]string) {
	for k, v := range src {
		a.db[k] = v
	}
}

func (a *Env) init() error {
	if a.Type == "" {
		return errors.New("env type is empty")
	}
	log.Debug("init %s", a.Type)
	files := []string{
		filepath.Join(a.cfg.LibDir, "env", "config.json"),
		filepath.Join(a.cfg.LibDir, "env", a.Type, "config.json"),
		filepath.Join(a.cfg.LibDir, "env", a.Type, a.OS+".json"),
	}
	for _, fn := range files {
		if fh, err := os.Open(fn); err != nil {
			log.Debug("%s", err)
			return err
		} else {
			log.Debug("load %s", fn)
			if err := a.load("", fh); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *Env) setRuntime() {
	a.OS = runtime.GOOS
	a.db["os"] = a.OS
	a.db["arch"] = runtime.GOARCH
}

func (a *Env) load(prefix string, fh io.ReadCloser) error {
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

func (a *Env) loadOS() error {
	if a.OS == "" {
		return errors.New("runtime OS is empty!?? =(")
	}
	files := []string{
		filepath.Join(a.cfg.LibDir, "os", "config.json"),
		filepath.Join(a.cfg.LibDir, "os", a.OS+".json"),
	}
	for _, fn := range files {
		if fh, err := os.Open(fn); err != nil {
			log.Debug("%s", err)
			return err
		} else {
			log.Debug("load %s", fn)
			if err := a.load("os", fh); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *Env) loadService() error {
	files := []string{
		filepath.Join(a.cfg.LibDir, "service", "config.json"),
		filepath.Join(a.cfg.LibDir, "service", a.Service, "config.json"),
		filepath.Join(a.cfg.LibDir, "service", a.Service, a.OS+".json"),
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
	return nil
}

func (a *Env) Get(opt string) string {
	return strings.TrimSpace(a.db[opt])
}

func (a *Env) GetAll(opt string) map[string]string {
	d := make(map[string]string)
	x := opt + "."
	for k, v := range a.db {
		if strings.HasPrefix(k, x) {
			n := strings.Replace(k, x, "", 1)
			d[n] = strings.TrimSpace(v)
		}
	}
	return d
}

func (a *Env) Update(opt, val string) error {
	_, ok := a.db[opt]
	if !ok {
		return errors.New("invalid option " + opt)
	}
	a.db[opt] = val
	return nil
}
