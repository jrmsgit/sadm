// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"github.com/jrmsdev/sadm/internal/log"
)

func (j *Jail) Start() error {
	log.Debug("%s", j.env.Name)
	if err := j.checkDestdir(); err != nil {
		return err
	}
	if err := j.mount(); err != nil {
		return err
	}
	return nil
}
