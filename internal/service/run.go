// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package service

import (
	"errors"
	"fmt"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
)

var sprintf = fmt.Sprintf

func Run(ctx *env.Env, action string) error {
	log.Debug("%s %s", action, ctx.Name)
	if ctx.Service == "" {
		return errors.New(sprintf("%s service is empty", ctx.Name))
	}
	// get service executable file
	cmd := ctx.Get("service.exec")
	if cmd == "" {
		return errors.New(sprintf("%s service exec is empty", ctx.Service))
	}
	log.Debug("%s exec %s", ctx.Service, cmd)
	if action == "create" {
		return create(ctx, cmd)
	} else if action == "check" {
		return check(ctx, cmd)
	} else if action == "start" {
		return start(ctx, cmd)
	} else if action == "stop" {
		return stop(ctx, cmd)
	}
	msg := sprintf("%s invalid service action %s", ctx.Name, action)
	log.Debug("%s", msg)
	return errors.New(msg)
}
