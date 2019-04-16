// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
	"github.com/jrmsdev/sadm/internal/utils/pkg"
)

func (j *Jail) Check(destdir, cmd string) error {
	// check jail destdir
	if !fs.Exists(destdir) {
		e := sprintf("%s dir not found", destdir)
		log.Debug("%s", e)
		return errors.New(e)
	}
	// check service packages/files deps
	if _, err := pkg.Check(j.args, cmd); err != nil {
		return err
	}
	return nil
}
