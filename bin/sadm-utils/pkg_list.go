// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"errors"
	"sort"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/pkg"
)

func pkgList(opt *args.Args, cmdargs []string) error {
	service := opt.Get("service")
	if service == "" {
		return errors.New("pkg.list no service name")
	}
	serviceExec := opt.Get("service.exec")
	if serviceExec == "" {
		return errors.New("pkg.list service exec is empty")
	}
	if len(cmdargs) != 0 {
		return errors.New(sprintf("pkg.list invalid args %v", cmdargs))
	}
	info, err := pkg.Which(opt, serviceExec)
	if err != nil {
		return err
	}
	log.Debug("pkg list %s", info.Pkg)
	if err := pkg.List(opt, info); err != nil {
		return err
	}
	return pkgReport(info)
}

func pkgReport(info *pkg.Info) error {
	log.Printf("Package: %s", info.Pkg)
	log.Printf("Deps (%d):", len(info.Deps))
	for _, dep := range info.Deps {
		log.Printf("  %s", dep.Pkg)
	}
	log.Printf("Files (%d):", len(info.Files))
	sort.Strings(info.Files)
	for _, fn := range info.Files {
		log.Printf("  %s", fn)
	}
	log.Printf("Files Prune (%d):", len(info.FilesPrune))
	sort.Strings(info.FilesPrune)
	for _, fn := range info.FilesPrune {
		log.Printf("  %s", fn)
	}
	return nil
}
