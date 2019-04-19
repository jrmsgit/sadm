// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/jrmsdev/sadm/internal/log"
)

func (e *Env) Dump(args []string) (string, error) {
	log.Debug("dump %v", args)
	largs := len(args)
	if largs > 0 {
		report := make([]string, 0)
		for _, p := range args {
			for k := range e.db {
				if strings.HasPrefix(k, p) {
					report = append(report, k)
				}
			}
		}
		sort.Strings(report)
		s := ""
		for _, k := range report {
			s = fmt.Sprintf("%s%s: %s\n", s, k, e.db[k])
		}
		return strings.TrimSpace(s), nil
	} else {
		if blob, err := json.MarshalIndent(e.db, "", "  "); err != nil {
			return "{}", err
		} else {
			return string(blob), nil
		}
	}
}
