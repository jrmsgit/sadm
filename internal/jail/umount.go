// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jail

import (
	"sort"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func (j *Jail) umount() error {
	log.Debug("%s", j.destdir)
	ls, err := fs.LsMount(j.destdir)
	if err != nil {
		return err
	}
	sort.Strings(ls)
	//~ log.Debug("mount list: %d %v", len(ls), ls)
	for idx := len(ls) - 1; idx >= 0; idx-- {
		mnt := ls[idx]
		log.Debug("umount %s", mnt)
		if err := fs.Umount(mnt); err != nil {
			return err
		}
	}
	return nil
}
