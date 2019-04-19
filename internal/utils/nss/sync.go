// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package nss

import (
	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
)

func Sync(ctx *env.Env, db string) error {
	destdir := ctx.Get("destdir")
	log.Debug("%s %s %s", ctx.Name, db, destdir)
	return nil
}
