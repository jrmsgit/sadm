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
	// service executable file
	cmd := j.args.Get("service.exec")
	if cmd == "" {
		return errors.New(sprintf("%s service exec is empty", j.args.Service))
	}
	log.Debug("%s exec %s", j.args.Service, cmd)
	// dispatch action
	if action == "check" {
		return j.Check(destdir, cmd)
	} else if action == "create" {
		return j.Create(destdir, cmd)
	} else if action == "start" {
		return j.Start(destdir)
	}
	// fail due to invalid action request
	err := errors.New(sprintf("invalid jail action %s", action))
	log.Debug("%s", err)
	return err
}
