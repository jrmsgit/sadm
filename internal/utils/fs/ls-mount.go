// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/log"
	//~ "github.com/jrmsdev/sadm/internal/utils"
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
	//~ if _, err := utils.Exec("/bin/mount", args...); err != nil {
		//~ return l, err
	//~ }
	return l, nil
}
