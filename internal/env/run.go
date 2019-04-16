// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/service"
)

var ValidAction = map[string]bool{
	"dump":   true,
	"check":  true,
	"create": true,
	"start":  true,
	"stop":   true,
}

func Run(e *Env, action string) error {
	log.Debug("%s %s %s", action, e.args.Type, e.name)
	if action == "dump" {
		if s, err := e.args.Dump(); err != nil {
			return err
		} else {
			log.Print(s)
			return nil
		}
	}
	if err := e.ctl.Dispatch(action); err != nil {
		return err
	}
	if action == "start" || action == "stop" {
		if err := service.Run(e.args, action); err != nil {
			return err
		}
	}
	return nil
}
