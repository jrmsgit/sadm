// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"path/filepath"
	"strings"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func (j *Jail) mount() error {
	log.Debug("%s %s", j.env.Name, j.destdir)
	d, err := filepath.Abs(j.destdir)
	if err != nil {
		log.Debug("%s", err)
		return err
	}
	for name, args := range j.env.GetAll("mount") {
		args = sprintf(args, d)
		log.Debug("%s mount %s", name, args)
		if err := fs.Mount(strings.Split(args, " ")); err != nil {
			return err
		}
	}
	return nil
}
