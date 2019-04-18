// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"encoding/json"

	"github.com/jrmsdev/sadm/internal/log"
)

func (a *Env) Dump() (string, error) {
	log.Debug("dump")
	if blob, err := json.MarshalIndent(a.db, "", "  "); err != nil {
		return "{}", err
	} else {
		return string(blob), nil
	}
}
