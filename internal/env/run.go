// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/log"
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
		return e.ctl.Check()
	}
	return errors.New("invalid action: " + action)
}
