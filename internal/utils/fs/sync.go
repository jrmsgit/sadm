// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"errors"
	"path/filepath"
	"sort"

	"github.com/jrmsdev/sadm/internal/log"
)

func Sync(destdir string, files ...string) error {
	log.Debug("sync %s %d files", destdir, len(files))
	sort.Strings(files)
	for _, fn := range files {
		var (
			err    error
			src    *Info
			srcdir *Info
		)
		src, err = newFileInfo(fn)
		if err != nil {
			return err
		}
		if src.IsDir() {
			if err := syncDir(destdir, src); err != nil {
				return err
			}
		} else {
			sdn := filepath.Dir(src.Filename())
			srcdir, err = newFileInfo(sdn)
			if err != nil {
				return err
			}
			if err := syncDir(destdir, srcdir); err != nil {
				return err
			}
			if err := syncFile(destdir, src); err != nil {
				return err
			}
		}
	}
	return nil
}

func syncDir(destdir string, src *Info) error {
	dst := destdir + src.Filename()
	if err := Mkdir(dst); err != nil {
		return err
	}
	if err := Chmod(dst, src.Mode()); err != nil {
		return err
	}
	return nil
}

func syncFile(destdir string, src *Info) error {
	if !src.Mode().IsRegular() {
		e := errors.New(sprintf("%s unsupported file type %s", src, src.Mode()))
		log.Debug("%s", e)
		return e
	}
	dst := destdir + src.Filename()
	if err := Copy(dst, src.Filename()); err != nil {
		return err
	}
	if err := Chmod(dst, src.Mode()); err != nil {
		return err
	}
	return nil
}
