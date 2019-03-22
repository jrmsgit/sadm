// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/service"
)

func Run(e *Env, action string) error {
	log.Debug("%s %s %s", e.name, action, e.Type())
	if action == "dump" {
		if s, err := e.Dump(); err != nil {
			return err
		} else {
			log.Print(s)
			return nil
		}
	} else if action == "check" {
		if err := e.ctl.Check(); err != nil {
			return err
		}
		return service.Check(e.args)
	}
	return errors.New("invalid action: " + action)
}
