// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

var sprintf = fmt.Sprintf

type Jail struct {
	env         *env.Env
	destdir     string
	serviceExec string
}

func New(opt *env.Env) (*Jail, error) {
	log.Debug("new %s", opt.Name)
	j := new(Jail)
	j.env = opt
	if err := j.setDefaults(); err != nil {
		return nil, err
	}
	if err := j.load(); err != nil {
		return nil, err
	}
	return j, nil
}

func (j *Jail) setDefaults() error {
	log.Debug("set defaults %s", j.env.Name)
	destdir := filepath.FromSlash(j.env.Get("destdir"))
	if destdir == "" {
		return errors.New("jail destdir is not set")
	}
	if err := j.env.Update("destdir", filepath.Join(destdir, j.env.Name)); err != nil {
		return err
	}
	return nil
}

func (j *Jail) load() error {
	// load destdir
	j.destdir = filepath.Clean(j.env.Get("destdir"))
	log.Debug("destdir %s", j.destdir)
	// load service exec
	j.serviceExec = strings.TrimSpace(j.env.Get("service.exec"))
	log.Debug("%s exec %s", j.env.Service, j.serviceExec)
	return nil
}

func (j *Jail) checkDestdir() error {
	if !fs.Exists(j.destdir) {
		err := errors.New(sprintf("%s dir not found", j.destdir))
		log.Debug("%s", err)
		return err
	}
	return nil
}

func (j *Jail) checkServiceExec() error {
	if j.serviceExec == "" {
		err := errors.New(sprintf("%s service exec is empty", j.env.Service))
		log.Debug("%s", err)
		return err
	}
	return nil
}
