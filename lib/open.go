// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package lib

import (
	"io"
	"os"
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/log"
)

func Open(filename string) (io.ReadCloser, error) {
	//~ log.Debug("%s", filename)
	fn := filepath.Join(prefix, filename)
	fh, err := os.Open(fn)
	if err != nil {
		if zipExists(filename) {
			return zipOpen(filename)
		}
		//~ log.Debug("%s", err)
		return nil, err
	}
	return fh, nil
}

func zipExists(filename string) bool {
	_, ok := storage[filename]
	if !ok {
		log.Debug("zip %s not found", filename)
	}
	return ok
}

func zipOpen(filename string) (io.ReadCloser, error) {
	//~ log.Debug("zip %s", filename)
	fh := storage[filename]
	return fh.Open()
}
