// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package service

import (
	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/tpl"
)

func configure(ctx *env.Env) error {
	log.Debug("%s", ctx.Name)
	if err := tpl.Make(ctx, "service"); err != nil {
		return err
	}
	log.Printf("service %s configured", ctx.Service)
	return nil
}
