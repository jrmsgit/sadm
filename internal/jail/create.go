// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func (j *Jail) Create() error {
	// create destdir
	if fs.Exists(j.destdir) {
		err := errors.New(sprintf("%s already exists", j.destdir))
		log.Debug("%s", err)
		return err
	}
	if err := fs.Mkdir(j.destdir); err != nil {
		return err
	}
	return nil
}
