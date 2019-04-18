// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils"
)

func Umount(mnt string) error {
	log.Debug("%s", mnt)
	args := []string{"-v", mnt}
	if out, err := utils.Exec("/bin/umount", args...); err != nil {
		return err
	} else {
		log.Print(out)
	}
	return nil
}
