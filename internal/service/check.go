// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package service

import (
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func Check(opt *args.Args) error {
	log.Debug(opt.Service)
	//~ s, _ := opt.Dump()
	//~ log.Print(s)
	return nil
}
