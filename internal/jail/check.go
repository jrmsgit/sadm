// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	//~ "github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func (j *Jail) Check() error {
	log.Debug("%s", j.opt.Get("service"))
	return nil
}
