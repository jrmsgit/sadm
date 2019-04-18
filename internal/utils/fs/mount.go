// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"strings"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils"
)

func Mount(args []string) error {
	log.Debug("%v", args)
	x := make([]string, 0)
	x = append(x, "-v")
	x = append(x, args...)
	if out, err := utils.Exec("/bin/mount", x...); err != nil {
		return err
	} else {
		log.Print(strings.TrimSpace(string(out)))
	}
	return nil
}
