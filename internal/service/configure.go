// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package service

import (
	"errors"
	"strings"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/tpl"
	"github.com/jrmsdev/sadm/internal/utils"
)

func configure(ctx *env.Env) error {
	log.Debug("%s", ctx.Name)
	if err := tpl.Make(ctx, "service"); err != nil {
		return err
	}
	if err := runHooks(ctx); err != nil {
		return err
	}
	log.Printf("service %s configured", ctx.Service)
	return nil
}

func runHooks(ctx *env.Env) error {
	log.Debug("run %s service hooks", ctx.Service)
	for name, hook := range ctx.GetAll("service.configure") {
		log.Debug("hook %s", name)
		hook = strings.TrimSpace(sprintf(hook, ctx.Get("destdir")))
		args := strings.Split(hook, " ")
		if len(args) < 2 {
			err := errors.New(sprintf("invalid service hook %s", name))
			log.Debug("%s", err)
			return err
		}
		if _, err := utils.Exec(args[0], args[1:]...); err != nil {
			return err
		} else {
			log.Printf("configure hook %s done", name)
		}
	}
	return nil
}
