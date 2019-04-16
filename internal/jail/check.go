// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/pkg"
)

func (j *Jail) Check() error {
	log.Debug("check %s", j.args.Env)
	if err := j.checkDestdir(); err != nil {
		return err
	}
	if _, err := pkg.Check(j.args, j.serviceExec); err != nil {
		return err
	}
	return nil
}
