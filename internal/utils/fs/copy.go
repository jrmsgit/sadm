// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"io"
	"os"

	"github.com/jrmsdev/sadm/internal/log"
)

func Copy(dst, src string) error {
	var (
		err error
		sfh io.ReadCloser
		dfh io.WriteCloser
	)
	sfh, err = os.Open(src)
	if err != nil {
		log.Debug("%s", err)
		return err
	}
	defer sfh.Close()
	flags := os.O_WRONLY | os.O_CREATE
	dfh, err = os.OpenFile(dst, flags, 0644)
	if err != nil {
		log.Debug("%s", err)
		return err
	}
	defer dfh.Close()
	_, err = io.Copy(dfh, sfh)
	if err != nil {
		log.Debug("%s", err)
		return err
	}
	log.Printf("copy %s -> %s", src, dst)
	return nil
}
