// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package service

import (
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func start(env *args.Args, cmd string) error {
	log.Debug("%s %s", env.Env, cmd)
	log.Printf("%s started %s", env.Env, cmd)
	return nil
}
