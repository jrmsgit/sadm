// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

type Manager interface {
	Which(string) (string, error)
}

func newManager(opt *args.Args) (Manager, error) {
	pkgman := opt.Get("os.pkg")
	if pkgman == "" {
		return nil, errors.New("os pkg manager not set")
	}
	pkgcmd := opt.Get("os.pkg.exec")
	if pkgcmd == "" {
		return nil, errors.New("os pkg exec not set")
	}
	log.Debug("new pkg manager %s (%s)", pkgman, pkgcmd)
	if pkgman == "dpkg" {
		return dpkgNew(pkgcmd), nil
	}
	return nil, errors.New("invalid os pkg manager "+pkgman)
}
