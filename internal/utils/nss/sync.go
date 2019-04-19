// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package nss

import (
	"path/filepath"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func Sync(ctx *env.Env, prefix string) error {
	destdir := filepath.Join(ctx.Get("destdir"), "etc")
	log.Debug("%s '%s' %s", ctx.Name, prefix, destdir)
	if err := fs.Mkdir(destdir); err != nil {
		return err
	}
	k := ""
	if prefix != "" {
		k = prefix + "."
	}
	k = k + "nss"
	for db, val := range ctx.GetAll(k) {
		dst := filepath.Join(destdir, db)
		log.Debug("dst %s", dst)
		log.Debug("val: %s", val)
	}
	return nil
}
