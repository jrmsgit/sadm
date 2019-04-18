// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func (j *Jail) umount() error {
	log.Debug("%s", j.destdir)
	ls, err := fs.LsMount(j.destdir)
	if err != nil {
		return err
	}
	for _, mnt := range ls {
		log.Debug("umount %s", mnt)
		if err := fs.Umount(mnt); err != nil {
			return err
		}
	}
	return nil
}
