// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"path/filepath"
	"strings"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils"
)

func LsMount(basedir string) ([]string, error) {
	log.Debug("%s", basedir)
	l := make([]string, 0)
	d, err := filepath.Abs(basedir)
	if err != nil {
		log.Debug("%s", err)
		return l, err
	}
	log.Debug("abs: %s", d)
	if out, err := utils.Exec("/bin/mount", "-l"); err != nil {
		return l, err
	} else {
		log.Print(strings.TrimSpace(string(out)))
	}
	return l, nil
}
