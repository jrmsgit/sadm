// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"
	//~ "os"
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
	// copy jail files
	err = j.copyFiles(info, destdir)
	if err != nil {
		return err
	}
	return nil
}

func (j *Jail) copyFiles(info *pkg.Info, destdir string) error {
	log.Debug("copy files %s %d %s", info.Pkg, len(info.Files), destdir)
	//~ for _, fn := range info.Files {
		//~ dst := destdir+fn
		//~ if err := fs.Copy(fn, dst)
	//~ }
	return nil
}
