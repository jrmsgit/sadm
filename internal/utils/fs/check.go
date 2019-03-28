// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"os"
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func Check(opt *args.Args, files ...string) error {
	if _, err := checkDiff(opt, files...); err != nil {
		return err
	}
	return nil
}

func checkDiff(opt *args.Args, files ...string) ([]string, error) {
	log.Debug("check diff %s: %d files", opt.Env, len(files))
	destdir := opt.Get("destdir")
	if destdir != "" {
		log.Debug("destdir %s", destdir)
	}
	l := make([]string, 0)
	for _, fn := range files {
		var (
			err error
			src os.FileInfo
			dst os.FileInfo
		)
		fn := filepath.Clean(fn)
		if fn == "." || fn == "/" {
			continue
		}
		src, err = os.Stat(fn)
		if err != nil {
			return nil, err
		}
		dst, err = os.Stat(destdir+fn)
		if err != nil {
			log.Warn(err)
			l = append(l, fn)
		} else {
			if diff(src, dst) {
				l = append(l, fn)
			}
		}
	}
	return l, nil
}
