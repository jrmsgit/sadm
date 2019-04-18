// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"strings"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils"
)

func Umount(mnt string) error {
	log.Debug("%s", mnt)
	if out, err := utils.Exec("/bin/umount", "-v", mnt); err != nil {
		return err
	} else {
		log.Print(strings.TrimSpace(string(out)))
	}
	return nil
}
