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

func (j *Jail) Create() error {
	// jail destdir
	destdir := filepath.Clean(j.args.Get("destdir"))
	if destdir == "" {
		return errors.New(sprintf("%s jail destdir is empty", j.args.Service))
	}
	log.Debug("destdir %s", destdir)
	e := "error"
	if s, err := os.Stat(destdir); err == nil {
		if s.IsDir() {
			e = sprintf("%s dir already exists", destdir)
		} else {
			e = sprintf("%s already exists", destdir)
		}
		log.Debug("%s", e)
		return errors.New(e)
	}
	if err := fs.Mkdir(destdir); err != nil {
		return err
	}
	// service executable file
	cmd := j.args.Get("service.exec")
	if cmd == "" {
		return errors.New(sprintf("%s service exec is empty", j.args.Service))
	}
	log.Debug("%s cmd %s", j.args.Service, cmd)
	// service pkg
	var (
		info *pkg.Info
		err  error
	)
	info, err = pkg.Which(j.args, cmd)
	if err != nil {
		return err
	}
	log.Debug("which %s: %s", cmd, info.Pkg)
	err = pkg.List(j.args, info)
	if err != nil {
		return err
	}
	// sync jail files
	err = fs.Sync(destdir, info.Files...)
	if err != nil {
		return err
	}
	return nil
}
