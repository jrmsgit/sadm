// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"strings"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils"
)

func LsMount(basedir string) ([]string, error) {
	log.Debug("%s", basedir)
	l := make([]string, 0)
	if out, err := utils.Exec("/bin/mount", "-l"); err != nil {
		return l, err
	} else {
		for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
			mnt := strings.Split(line, " ")[2]
			if strings.HasPrefix(mnt, basedir) {
				l = append(l, mnt)
			}
		}
	}
	return l, nil
}
