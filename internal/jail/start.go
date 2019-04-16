// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"github.com/jrmsdev/sadm/internal/log"
)

func (j *Jail) Start() error {
	log.Debug("start %s", j.args.Env)
	if err := j.checkDestdir(); err != nil {
		return err
	}
	return nil
}
