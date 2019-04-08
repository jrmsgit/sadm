// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/log"
)

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
	}
	log.Debug("invalid action %s", action)
	return errors.New("run invalid action " + action)
}
