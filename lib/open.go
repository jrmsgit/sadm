// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package lib

import (
	"io"

	"github.com/jrmsdev/gojc/fs"
)

func Open(filename string) (io.ReadCloser, error) {
	return fs.Open(prefix, filename)
}
