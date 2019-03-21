// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

type Jail struct {
}

func New(opt *args.Args) (*Jail, error) {
	log.Debug("new")
	return &Jail{}, nil
}

func (j *Jail) Check() error {
	log.Debug("check")
	return nil
}
