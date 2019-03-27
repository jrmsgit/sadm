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
	cmd := j.args.Get("service.exec")
	if cmd == "" {
		return errors.New(sprintf("%s exec is empty", j.args.Service))
	}
	log.Debug("%s cmd %s", j.args.Service, cmd)
	if info, err := ldd.List(j.args, cmd); err != nil {
		return err
	} else {
		for _, fn := range info.Files {
			log.Print(fn)
		}
	}
	return nil
}
