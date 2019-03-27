// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"

	//~ "github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/ldd"
)

func (j *Jail) Check() error {
	cmd := j.opt.Get("service.exec")
	if cmd == "" {
		return errors.New(sprintf("%s exec is empty", j.opt.Service))
	}
	log.Debug("%s cmd %s", j.opt.Service, cmd)
	if fl, err := ldd.List(cmd); err != nil {
		return err
	} else {
		for _, fn := range fl {
			log.Print(fn)
		}
	}
	return nil
}
