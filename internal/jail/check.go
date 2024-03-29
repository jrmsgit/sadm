// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"github.com/jrmsdev/sadm/internal/log"
)

func (j *Jail) Check() error {
	log.Debug("%s", j.env.Name)
	if err := j.checkDestdir(); err != nil {
		return err
	}
	log.Printf("jail %s dir found", j.destdir)
	return nil
}
