// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/jrmsdev/sadm/internal/cfg"
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/jail"
	"github.com/jrmsdev/sadm/internal/log"
)

var sprintf = fmt.Sprintf

type Env struct {
	name string
	args *args.Args
	ctl  Manager
	//~ cfg  *cfg.Cfg
}

func New(config *cfg.Cfg, name string, src io.ReadCloser) (*Env, error) {
	log.Debug("new %s", name)
	environ := new(Env)
	environ.name = name
	//~ environ.cfg = config
	defer src.Close()
	// parse args
	if blob, err := ioutil.ReadAll(src); err != nil {
		return nil, err
	} else {
		a := make(map[string]string)
		if err := json.Unmarshal(blob, &a); err != nil {
			return nil, err
		}
		environ.args, err = args.New(config, a)
		if err != nil {
			return nil, err
		}
	}
	// env manager
	if err := newManager(environ); err != nil {
		return nil, err
	}
	// env service
	if environ.args.Service == "" {
		return nil, errors.New(sprintf("%s: service definition is empty", name))
	}
	//~ log.Debug("%#v", environ)
	return environ, nil
}

func newManager(e *Env) error {
	log.Debug("new manager %s", e.name)
	typ := e.Type()
	if typ == "" {
		return errors.New(sprintf("%s: type definition is empty", e.name))
	}
	var err error
	if typ == "jail" {
		e.ctl, err = jail.New(e.args)
		if err != nil {
			return err
		}
	} else {
		return errors.New(sprintf("%s: invalid type %s", e.name, typ))
	}
	return nil
}

func (e *Env) Dump() (string, error) {
	return e.args.Dump()
}

func (e *Env) String() string {
	return e.name
}

func (e *Env) Type() string {
	return e.args.Get("type")
}
