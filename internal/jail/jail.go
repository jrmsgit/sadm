// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

var sprintf = fmt.Sprintf

type Jail struct {
	args *args.Args
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
