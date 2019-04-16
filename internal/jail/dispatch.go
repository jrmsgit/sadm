// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func (j *Jail) Dispatch(action string) error {
	// check jail destdir
	destdir := filepath.Clean(j.args.Get("destdir"))
	if destdir == "" {
		err := errors.New(sprintf("%s jail destdir is empty", j.args.Service))
		log.Debug("%s", err)
		return err
	}
	log.Debug("destdir %s", destdir)
	if !fs.Exists(destdir) {
		err := errors.New(sprintf("%s dir not found", destdir))
		log.Debug("%s", err)
		return err
	}
	if action == "check" {
		return j.Check()
	} else if action == "create" {
		return j.Create()
	} else if action == "start" {
		return j.Start()
	}
	err := errors.New(sprintf("invalid jail action %s", action))
	log.Debug("%s", err)
	return err
}
