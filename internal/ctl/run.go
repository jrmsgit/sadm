// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ctl

import (
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/service"
)

var ValidAction = map[string]bool{
	"dump":      true,
	"check":     true,
	"create":    true,
	"configure": true,
	"start":     true,
	"stop":      true,
}

var serviceFirst = map[string]bool{
	"stop": true,
}

func Run(x *Ctl, action string) error {
	log.Debug("%s %s %s", action, x.env.Type, x.name)
	if action == "dump" {
		if s, err := x.env.Dump(); err != nil {
			return err
		} else {
			log.Print(s)
			return nil
		}
	}
	if serviceFirst[action] {
		if err := service.Run(x.env, action); err != nil {
			return err
		}
	}
	if err := x.man.Dispatch(action); err != nil {
		return err
	}
	if !serviceFirst[action] {
		if err := service.Run(x.env, action); err != nil {
			return err
		}
	}
	return nil
}
