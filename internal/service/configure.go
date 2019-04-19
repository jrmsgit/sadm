// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package service

import (
	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
)

func configure(ctx *env.Env) error {
	log.Debug("%s", ctx.Name)
	log.Printf("service %s configured", ctx.Service)
	return nil
}
