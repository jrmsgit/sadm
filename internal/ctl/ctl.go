// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ctl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/cfg"
	"github.com/jrmsdev/sadm/internal/jail"
	"github.com/jrmsdev/sadm/internal/log"
)

var sprintf = fmt.Sprintf

type Ctl struct {
	name string
	env  *env.Env
	man  Manager
	//~ cfg  *cfg.Cfg
}

func New(config *cfg.Cfg, name string, src io.ReadCloser) (*Ctl, error) {
	log.Debug("new %s", name)
	x := new(Ctl)
	x.name = name
	//~ x.cfg = config
	defer src.Close()
	// parse args
	if blob, err := ioutil.ReadAll(src); err != nil {
		return nil, err
	} else {
		a := make(map[string]string)
		if err := json.Unmarshal(blob, &a); err != nil {
			return nil, err
		}
		x.env, err = env.New(config, name, a)
		if err != nil {
			return nil, err
		}
	}
	// env xager
	if err := newManager(x); err != nil {
		return nil, err
	}
	// env service
	if x.env.Service == "" {
		return nil, errors.New(sprintf("%s: service definition is empty", name))
	}
	//~ log.Debug("%#v", x)
	return x, nil
}

func newManager(x *Ctl) error {
	typ := x.env.Type
	if typ == "" {
		return errors.New(sprintf("%s type definition is empty", x.name))
	}
	log.Debug("%s %s manager", x.name, typ)
	var err error
	if typ == "jail" {
		x.man, err = jail.New(x.env)
		if err != nil {
			return err
		}
	} else {
		return errors.New(sprintf("%s: invalid type %s", x.name, typ))
	}
	return nil
}
