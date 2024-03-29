// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"errors"
	"os"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
)

func Check(opt *env.Env, files ...string) error {
	if l, err := checkDiff(opt, files...); err != nil {
		return err
	} else {
		llen := len(l)
		if llen > 0 {
			log.Debug("fs diff found: %s %d", opt.Service, llen)
			return errors.New(sprintf("%s fs diff found: %d files", opt.Service, llen))
		}
	}
	log.Printf("%s: %d files checked", opt.Name, len(files))
	return nil
}

func checkDiff(opt *env.Env, files ...string) ([]string, error) {
	log.Debug("check diff %s: %d files", opt.Name, len(files))
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
		src, err = os.Stat(fn)
		if err != nil {
			return nil, err
		}
		dstfn := destdir + fn
		dst, err = os.Stat(dstfn)
		if err != nil {
			log.Warn(err)
			l = append(l, fn)
		} else {
			if diff(newInfo(src, fn), newInfo(dst, dstfn)) {
				l = append(l, fn)
			}
		}
	}
	return l, nil
}
