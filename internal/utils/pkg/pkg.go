// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	"errors"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
)

var depDone = errors.New("depdone")

type Info struct {
	Pkg        string
	Deps       []*Info
	Files      []string
	FilesPrune []string
}

func (i *Info) String() string {
	return i.Pkg
}

type Manager interface {
	Which(*Info, string) error
	Depends(*Info, string) error
	List(*Info, string) error
}

func newManager(opt *env.Env) (Manager, error) {
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
		return dpkgNew(opt), nil
	}
	return nil, errors.New("invalid os pkg manager " + pkgman)
}
