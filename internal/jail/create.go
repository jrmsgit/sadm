// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
	"github.com/jrmsdev/sadm/internal/utils/pkg"
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
	// service pkg
	var (
		info *pkg.Info
		err  error
	)
	if err := j.checkServiceExec(); err != nil {
		return err
	}
	info, err = pkg.Which(j.args, j.serviceExec)
	if err != nil {
		return err
	}
	log.Debug("which %s: %s", j.serviceExec, info.Pkg)
	err = pkg.List(j.args, info)
	if err != nil {
		return err
	}
	// sync jail files
	err = fs.Sync(j.destdir, info.Files...)
	if err != nil {
		return err
	}
	return nil
}
