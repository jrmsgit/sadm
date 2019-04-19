// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package tpl

import (
	"io/ioutil"
	"path/filepath"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func Make(ctx *env.Env, prefix string) error {
	destdir := ctx.Get("destdir")
	log.Debug("%s env prefix='%s' destdir='%s'", ctx.Name, prefix, destdir)
	for name, filename := range ctx.GetAll(prefix + ".template") {
		if err := mktpl(ctx, destdir, prefix + "." + name, filename); err != nil {
			return err
		}
	}
	return nil
}

func mkdst(fn string) error {
	dst := filepath.Dir(fn)
	return fs.Mkdir(dst)
}

func mktpl(ctx *env.Env, destdir, name, filename string) error {
	fn := ctx.TplFile(filename)
	dst := filepath.Join(destdir, filename)
	log.Debug("tpl %s", fn)
	log.Debug("dst %s", dst)
	if err := mkdst(dst); err != nil {
		return err
	}
	blob, err := parseFile(fn, name, ctx.TplData())
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, blob, 0644)
	if err != nil {
		log.Debug("%s", err)
		return err
	}
	return nil
}
