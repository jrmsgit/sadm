// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ldd

import (
	"path/filepath"
	"strings"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils"
)

func List(opt *args.Args, filename string) (*Info, error) {
	x := opt.Get("ldd")
	if x == "" {
		x = "ldd"
	}
	log.Debug("%s %s", x, filename)
	l := newInfo(filename)
	if out, err := utils.Exec(x, filename); err != nil {
		log.Debug("%s", err)
		return nil, err
	} else {
		if err := parse(&l.Files, string(out)); err != nil {
			return nil, err
		}
	}
	return l, nil
}

func parse(dst *[]string, src string) error {
	log.Debug("parse")
	for _, line := range strings.Split(src, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		//~ log.Print(line)
		items := strings.Split(line, " => ")
		ilen := len(items)
		tok0 := strings.TrimSpace(items[0])
		if ilen == 1 {
			fn := strings.Split(tok0, " ")[0]
			checkFile(dst, fn)
		} else {
			tok1 := strings.TrimSpace(items[1])
			fn := strings.Split(tok1, " ")[0]
			checkFile(dst, fn)
		}
	}
	return nil
}

func checkFile(dst *[]string, fn string) {
	log.Debug("check file %s", fn)
	if !filepath.IsAbs(fn) {
		log.Warnf("ldd ignore %s", fn)
		return
	}
	*dst = append(*dst, fn)
}
