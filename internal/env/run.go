// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"github.com/jrmsdev/sadm/internal/log"
)

func Run(e *Env, action string) error {
	log.Debug("%s %s %s", e, action, e.Type())
	return nil
}
