// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
	"github.com/jrmsdev/sadm/internal/utils/pkg"
)

func (j *Jail) Check() error {
	cmd := j.args.Get("service.exec")
	if cmd == "" {
		return errors.New(sprintf("%s service exec is empty", j.args.Service))
	}
	log.Debug("%s cmd %s", j.args.Service, cmd)
	// service pkg check
	if info, err := pkg.Check(j.args, cmd); err != nil {
		return err
	} else {
		// jail files check
		if err := j.checkFiles(info); err != nil {
			return err
		}
	}
	return nil
}

func (j *Jail) checkFiles(info *pkg.Info) error {
	destdir := j.args.Get("destdir")
	if destdir == "" {
		return errors.New(sprintf("%s jail destdir is empty", j.args.Service))
	}
	log.Debug("check files %s destdir=%s", info.Pkg, destdir)
	return fs.Check(j.args, info.Files...)
}
