// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package service

import (
	"errors"
	"fmt"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

var sprintf = fmt.Sprintf

func Run(action string, env *args.Args) error {
	log.Debug("%s %s", action, env.Env)
	if env.Service == "" {
		return errors.New(sprintf("%s service is empty", env.Env))
	}
	// get service executable file
	cmd := env.Get("service.exec")
	if cmd == "" {
		return errors.New(sprintf("%s service exec is empty", env.Service))
	}
	log.Debug("%s exec %s", env.Service, cmd)
	if action == "start" {
		return start(env, cmd)
	//~ } else if action == "stop" {
		//~ return stop(env)
	}
	msg := sprintf("%s invalid service action %s", env.Env, action)
	log.Debug("%s", msg)
	return errors.New(msg)
}
