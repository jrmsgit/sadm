// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
	"github.com/jrmsdev/sadm/internal/utils/pkg"
)

func (j *Jail) Check() error {
	// service executable file
	cmd := j.args.Get("service.exec")
	if cmd == "" {
		return errors.New(sprintf("%s service exec is empty", j.args.Service))
	}
	log.Debug("%s cmd %s", j.args.Service, cmd)
	// jail destdir
	destdir := filepath.Clean(j.args.Get("destdir"))
	if destdir == "" {
		return errors.New(sprintf("%s jail destdir is empty", j.args.Service))
	}
	log.Debug("destdir %s", destdir)
	if _, err := os.Stat(destdir); err != nil {
		log.Debug("%s", err)
		return err
	}
	// service pkg
	if info, err := pkg.Check(j.args, cmd); err != nil {
		return err
	} else {
		// jail files
		if err := j.checkFiles(info); err != nil {
			return err
		}
	}
	return nil
}

func (j *Jail) checkFiles(info *pkg.Info) error {
	log.Debug("check files: %s", info.Pkg)
	return fs.Check(j.args, info.Files...)
}
