// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package service

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func start(env *args.Args) error {
	log.Debug(env.Env)
	if env.Service == "" {
		return errors.New(sprintf("%s service is empty", env.Env))
	}
	// get service executable file
	cmd := env.Get("service.exec")
	if cmd == "" {
		return errors.New(sprintf("%s service exec is empty", env.Service))
	}
	log.Debug("%s exec %s", env.Service, cmd)
	log.Printf("%s %s started", env.Env, cmd)
	return nil
}
