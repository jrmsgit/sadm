// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"strings"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils"
)

func Mkdir(dirname string) error {
	out, err := utils.Exec("mkdir", "-vp", dirname)
	if err != nil {
		return err
	}
	outs := strings.TrimSpace(string(out))
	if outs != "" {
		log.Print(outs)
	}
	return nil
}
