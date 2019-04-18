// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package service

import (
	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
)

func start(ctx *env.Env, cmd string) error {
	log.Debug("%s %s", ctx.Name, cmd)
	log.Printf("%s started %s", ctx.Name, cmd)
	return nil
}
