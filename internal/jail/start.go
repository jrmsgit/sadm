// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func (j *Jail) Start() error {
	// check jail destdir
	destdir := filepath.Clean(j.args.Get("destdir"))
	if destdir == "" {
		e := sprintf("%s jail destdir is empty", j.args.Service)
		log.Debug("%s", e)
		return errors.New(e)
	}
	log.Debug("destdir %s", destdir)
	if !fs.Exists(destdir) {
		e := sprintf("%s dir not found", destdir)
		log.Debug("%s", e)
		return errors.New(e)
	}
	return nil
}
