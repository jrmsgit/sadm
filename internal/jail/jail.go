// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

type Jail struct {
	opt *args.Args
}

func New(opt *args.Args) (*Jail, error) {
	log.Debug("new")
	return &Jail{opt}, nil
}
