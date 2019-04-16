// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

var sprintf = fmt.Sprintf

type Jail struct {
	args        *args.Args
	destdir     string
	serviceExec string
}

func New(opt *args.Args) (*Jail, error) {
	log.Debug("new %s", opt.Env)
	j := new(Jail)
	j.args = opt
	if err := j.setDefaults(); err != nil {
		return nil, err
	}
	if err := j.load(); err != nil {
		return nil, err
	}
	return j, nil
}

func (j *Jail) setDefaults() error {
	log.Debug("set defaults %s", j.args.Env)
	destdir := filepath.FromSlash(j.args.Get("destdir"))
	if destdir == "" {
		return errors.New("jail destdir is not set")
	}
	if err := j.args.Update("destdir", filepath.Join(destdir, j.args.Env)); err != nil {
		return err
	}
	return nil
}

func (j *Jail) load() error {
	// load destdir
	j.destdir = filepath.Clean(j.args.Get("destdir"))
	log.Debug("destdir %s", j.destdir)
	// load service exec
	j.serviceExec = strings.TrimSpace(j.args.Get("service.exec"))
	log.Debug("%s exec %s", j.args.Service, j.serviceExec)
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
		err := errors.New(sprintf("%s service exec is empty", j.args.Service))
		log.Debug("%s", err)
		return err
	}
	return nil
}
