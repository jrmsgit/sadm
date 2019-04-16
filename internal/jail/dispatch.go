// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/log"
)

func (j *Jail) Dispatch(action string) error {
	log.Debug("dispatch %s", j.args.Env)
	// dispatch action
	if action == "check" {
		return j.Check()
	} else if action == "create" {
		return j.Create()
	} else if action == "start" {
		return j.Start()
	}
	// fail due to invalid action request
	err := errors.New(sprintf("invalid jail action %s", action))
	log.Debug("%s", err)
	return err
}
