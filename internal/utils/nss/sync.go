// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package nss

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils"
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
	dbs := make(map[string]map[string]bool)
	for _, p := range syncPrefix {
		for db, val := range ctx.GetAll(p) {
			val = strings.TrimSpace(val)
			if val == "" {
				continue
			}
			_, ok := dbs[db]
			if !ok {
				dbs[db] = make(map[string]bool)
			}
			for _, k := range strings.Split(val, " ") {
				k = strings.TrimSpace(k)
				if k != "" {
					dbs[db][k] = true
				}
			}
		}
	}
	for db := range dbs {
		l := make([]string, 0)
		for k := range dbs[db] {
			l = append(l, k)
		}
		sort.Strings(l)
		llen := len(l)
		if llen > 0 {
			dst := filepath.Join(destdir, db)
			if err := syncDB(dst, db, l); err != nil {
				return err
			} else {
				log.Printf("nss %s sync done", db)
			}
		}
	}
	return nil
}

func syncDB(dst, db string, keys []string) error {
	log.Debug("sync db %s %v %s", db, keys, dst)
	args := make([]string, 0)
	args = append(args, db)
	args = append(args, keys...)
	out, err := utils.Exec("/usr/bin/getent", args...)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(dst, out, 0644); err != nil {
		log.Debug("%s", err)
		return err
	}
	return nil
}
