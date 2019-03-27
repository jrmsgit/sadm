// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

type Jail struct {
	opt *args.Args
}

func New(opt *args.Args) (*Jail, error) {
	log.Debug("new %s", opt.Env)
	j := &Jail{opt}
	if err := j.setDefaults(); err != nil {
		return nil, err
	}
	return j, nil
}

func (j *Jail) setDefaults() error {
	log.Debug("set defaults %s", j.opt.Env)
	destdir := filepath.FromSlash(j.opt.Get("destdir"))
	if destdir == "" {
		return errors.New("jail destdir is not set")
	}
	if err := j.opt.Update("destdir", filepath.Join(destdir, j.opt.Env)); err != nil {
		return err
	}
	return nil
}
