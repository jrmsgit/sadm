// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/nss"
)

func (j *Jail) Configure() error {
	log.Debug("%s", j.env.Name)
	if err := j.checkDestdir(); err != nil {
		return err
	}
	// sync nss databases
	if err := nss.Sync(j.env); err != nil {
		return err
	}
	log.Printf("jail %s configured", j.destdir)
	return nil
}
