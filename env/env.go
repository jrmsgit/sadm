// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/jrmsdev/sadm/internal/cfg"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/lib"
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
	e := new(Env)
	e.cfg = config
	e.db = make(map[string]string)
	e.Name = name
	e.setRuntime()
	e.Type = strings.TrimSpace(src["type"])
	if err := e.init(); err != nil {
		return nil, err
	}
	if err := e.loadOS(); err != nil {
		return nil, err
	}
	e.Service = strings.TrimSpace(src["service"])
	if e.Service != "" {
		if err := e.loadService(); err != nil {
			return nil, err
		}
	}
	if err := e.source(src); err != nil {
		return nil, err
	}
	//~ log.Debug("new %#v", a)
	return e, nil
}

func (e *Env) source(src map[string]string) error {
	log.Debug("source config")
	var err error
	for k, v := range src {
		err = e.Update(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Env) init() error {
	if e.Type == "" {
		return errors.New("env type is empty")
	}
	if err := lib.Init(); err != nil {
		return err
	}
	log.Debug("init %s", e.Type)
	files := []string{
		filepath.Join("env", "config.json"),
		filepath.Join("env", e.Type, "config.json"),
		filepath.Join("env", e.Type, e.OS+".json"),
	}
	for _, fn := range files {
		if err := e.loadFile(fn, ""); err != nil {
			return err
		}
	}
	return nil
}

func (e *Env) setRuntime() {
	e.OS = runtime.GOOS
	e.db["os"] = e.OS
	e.db["arch"] = runtime.GOARCH
}

func (e *Env) loadFile(filename, prefix string) error {
	fh, err := lib.Open(filename)
	if err != nil {
		errOk := false
		fn := filepath.Base(filename)
		if fn != "config.json" {
			errOk = true
		}
		if errOk {
			log.Debug("ignore error: %s", err)
			return nil
		}
		log.Debug("%s", err)
		return err
	}
	log.Debug("load %s", filename)
	return e.load(filename, prefix, fh)
}

func (e *Env) load(filename, prefix string, fh io.ReadCloser) error {
	defer fh.Close()
	src := make(map[string]string)
	if blob, err := ioutil.ReadAll(fh); err != nil {
		log.Debug("%s", err)
		return err
	} else {
		if err := json.Unmarshal(blob, &src); err != nil {
			e := errors.New("lib unmarshal error " + filename + ": " + err.Error())
			log.Debug("%s", e)
			return e
		} else {
			if prefix != "" {
				prefix = prefix + "."
			}
			for opt, val := range src {
				e.db[prefix+opt] = val
			}
			//~ log.Debug("%s loaded %#v", prefix, e.db)
		}
	}
	return nil
}

func (e *Env) loadOS() error {
	if e.OS == "" {
		return errors.New("runtime OS is empty!?? =(")
	}
	files := []string{
		filepath.Join("os", "config.json"),
		filepath.Join("os", e.OS+".json"),
	}
	for _, fn := range files {
		if err := e.loadFile(fn, "os"); err != nil {
			return err
		}
	}
	return nil
}

func (e *Env) loadService() error {
	files := []string{
		filepath.Join("service", "config.json"),
		filepath.Join("service", e.Service, "config.json"),
		filepath.Join("service", e.Service, e.OS+".json"),
	}
	for _, fn := range files {
		if err := e.loadFile(fn, "service"); err != nil {
			return err
		}
	}
	return nil
}

func (e *Env) Get(opt string) string {
	return strings.TrimSpace(e.db[opt])
}

func (e *Env) GetAll(opt string) map[string]string {
	d := make(map[string]string)
	x := opt + "."
	for k, v := range e.db {
		if strings.HasPrefix(k, x) {
			n := strings.Replace(k, x, "", 1)
			d[n] = strings.TrimSpace(v)
		}
	}
	return d
}

func (e *Env) Update(opt, val string) error {
	_, ok := e.db[opt]
	if !ok {
		err := errors.New("invalid option " + opt)
		log.Debug("%s", err)
		return err
	}
	e.db[opt] = val
	return nil
}
