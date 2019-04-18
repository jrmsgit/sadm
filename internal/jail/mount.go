// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"github.com/jrmsdev/sadm/internal/log"
)

func (j *Jail) mount() error {
	log.Debug("%s", j.env.Name)
	return nil
}
