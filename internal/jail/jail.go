// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

var sprintf = fmt.Sprintf

type Jail struct {
	env     *env.Env
	destdir string
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
	var err error
	// destdir
	destdir := filepath.FromSlash(j.env.Get("destdir"))
	if destdir == "" {
		return errors.New("jail destdir is not set")
	}
	destdir, err = filepath.Abs(destdir)
	if err != nil {
		log.Debug("%s", err)
		return err
	}
	if err := j.env.Update("destdir", filepath.Join(destdir, j.env.Name)); err != nil {
		return err
	}
	return nil
}

func (j *Jail) load() error {
	// destdir
	j.destdir = j.env.Get("destdir")
	log.Debug("destdir %s", j.destdir)
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
