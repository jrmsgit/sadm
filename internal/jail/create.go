// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
	"github.com/jrmsdev/sadm/internal/utils/pkg"
)

func (j *Jail) Create(destdir, cmd string) error {
	// create destdir
	if err := fs.Mkdir(destdir); err != nil {
		return err
	}
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
