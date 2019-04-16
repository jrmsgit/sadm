// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func (j *Jail) Start(destdir string) error {
	// check jail destdir
	if !fs.Exists(destdir) {
		e := sprintf("%s dir not found", destdir)
		log.Debug("%s", e)
		return errors.New(e)
	}
	return nil
}
