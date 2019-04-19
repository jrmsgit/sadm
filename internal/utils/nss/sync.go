// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package nss

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

var syncPrefix = []string{
	"nss",
	"service.nss",
}

func Sync(ctx *env.Env) error {
	destdir := filepath.Join(ctx.Get("destdir"), "etc")
	log.Debug("%s %s", ctx.Name, destdir)
	if err := fs.Mkdir(destdir); err != nil {
		return err
	}
	for _, p := range syncPrefix {
		for db, val := range ctx.GetAll(p) {
			val = strings.TrimSpace(val)
			if val == "" {
				continue
			}
			keys := make(map[string]bool)
			dst := filepath.Join(destdir, db)
			for _, k := range strings.Split(val, " ") {
				k = strings.TrimSpace(k)
				if k != "" {
					keys[k] = true
				}
			}
			l := make([]string, 0)
			for k, _ := range keys {
				l = append(l, k)
			}
			sort.Strings(l)
			llen := len(l)
			if llen > 0 {
				syncDB(dst, db, l)
			}
		}
	}
	return nil
}

func syncDB(dst, db string, keys []string) error {
	log.Debug("sync db %s %v %s", db, keys, dst)
	return nil
}
