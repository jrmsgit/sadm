// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"hash/adler32"
	"io/ioutil"

	"github.com/jrmsdev/sadm/internal/log"
)

func checksum(filename string) string {
	blob, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Warn(err)
		return "ERROR:" + filename
	}
	sum := adler32.Checksum(blob)
	//~ log.Debug("checksum %s %d", filename, sum)
	return sprintf("%d", sum)
}
