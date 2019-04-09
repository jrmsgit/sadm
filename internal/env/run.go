// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"errors"

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
	} else if action == "check" {
		return e.ctl.Check()
	} else if action == "create" {
		return e.ctl.Create()
	} else if action == "start" {
		if err := e.ctl.Start(); err != nil {
			return err
		}
		return service.Start(e.args)
	}
	log.Debug("invalid action %s", action)
	return errors.New("run invalid action " + action)
}
