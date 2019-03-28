// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"os"

	"github.com/jrmsdev/sadm/internal/log"
)

func diff(src, dst os.FileInfo) bool {
	log.Debug("diff %s %s", src, dst)
	return false
}
