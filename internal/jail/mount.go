// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"strings"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func (j *Jail) mount() error {
	log.Debug("%s %s", j.env.Name, j.destdir)
	for name, args := range j.env.GetAll("mount") {
		text := sprintf(args, j.destdir)
		log.Debug("%s mount %s", name, text)
		err := fs.Mount(strings.Split(text, " ")...)
		if err != nil {
			return err
		}
	}
	return nil
}
