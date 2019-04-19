// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"encoding/json"
	"strings"

	"github.com/jrmsdev/sadm/internal/log"
)

func (e *Env) Dump(args []string) (string, error) {
	log.Debug("dump %v", args)
	largs := len(args)
	var report map[string]string
	if largs > 0 {
		report = make(map[string]string)
		for _, p := range args {
			for k, v := range e.db {
				if strings.HasPrefix(k, p) {
					report[k] = v
				}
			}
		}
	}
	if report == nil {
		report = e.db
	}
	if blob, err := json.MarshalIndent(report, "", "  "); err != nil {
		return "{}", err
	} else {
		return string(blob), nil
	}
}
