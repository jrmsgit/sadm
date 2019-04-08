// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
	"github.com/jrmsdev/sadm/internal/utils/pkg"
)

func (j *Jail) Check() error {
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
	// get service executable file
	cmd := j.args.Get("service.exec")
	if cmd == "" {
		return errors.New(sprintf("%s service exec is empty", j.args.Service))
	}
	log.Debug("%s cmd %s", j.args.Service, cmd)
	// check service packages/files deps
	if _, err := pkg.Check(j.args, cmd); err != nil {
		return err
	}
	return nil
}
