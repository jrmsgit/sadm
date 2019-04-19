// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package service

import (
	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/pkg"
)

func check(ctx *env.Env, cmd string) error {
	log.Debug("check %s", ctx.Name)
	if _, err := pkg.Check(ctx, cmd); err != nil {
		return err
	}
	return nil
}
