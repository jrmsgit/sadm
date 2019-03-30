// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"errors"
	"strings"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/pkg"
)

func pkgList(opt *args.Args, cmdargs []string) error {
	if len(cmdargs) != 1 {
		return errors.New(sprintf("pkg.list invalid args %v", cmdargs))
	}
	pkgname := strings.TrimSpace(cmdargs[0])
	if pkgname == "" {
		return errors.New("pkg.list no package name")
	}
	log.Debug("pkg list %s", pkgname)
	info := &pkg.Info{}
	info.Pkg = pkgname
	if err := pkg.List(opt, info); err != nil {
		return err
	}
	return nil
}
