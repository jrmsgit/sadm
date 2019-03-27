// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	//~ "errors"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

type Manager interface {
	Which(string) (string, error)
}

func Check(opt *args.Args, filename string) error {
	log.Debug("check %s", filename)
	if pkgname, err := which(filename); err != nil {
		return err
	} else {
		log.Debug("%s provided by %s", filename, pkgname)
	}
	return nil
}

func which(filename string) (string, error) {
	n := ""
	return n, nil
}
