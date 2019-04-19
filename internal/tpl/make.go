// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package tpl

import (
	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
)

func Make(ctx *env.Env, prefix string) error {
	log.Debug("%s env prefix='%s'", ctx.Name, prefix)
	for name, filename := range ctx.GetAll(prefix + ".template") {
		log.Print(name, filename)
	}
	return nil
}
