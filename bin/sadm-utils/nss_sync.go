// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"errors"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/nss"
)

func nssSync(ctx *env.Env, cmdargs []string) error {
	log.Debug("%v", cmdargs)
	if len(cmdargs) < 2 {
		return errors.New("invalid args")
	}
	destdir := cmdargs[0]
	ctx.Update("destdir", destdir)
	db := cmdargs[1]
	return nss.Sync(ctx, db)
}
