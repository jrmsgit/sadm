// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package service

import (
	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
	"github.com/jrmsdev/sadm/internal/utils/pkg"
)

func create(ctx *env.Env, cmd string) error {
	log.Debug("%s %s", ctx.Name, cmd)
	// service pkg
	var (
		info *pkg.Info
		err  error
	)
	info, err = pkg.Which(ctx, cmd)
	if err != nil {
		return err
	}
	log.Debug("which %s: %s", cmd, info.Pkg)
	err = pkg.List(ctx, info)
	if err != nil {
		return err
	}
	// sync jail files
	err = fs.Sync(ctx.Get("destdir"), info.Files...)
	if err != nil {
		return err
	}
	log.Printf("%s created %s", ctx.Name, cmd)
	return nil
}
