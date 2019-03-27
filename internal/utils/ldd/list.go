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
	if err := list(l, x, filename); err != nil {
		return nil, err
	}
	count := 0
	for fn, ok := range l.Files {
		if ok {
			count += 1
			list(l, x, fn)
		} else {
			log.Warnf("ldd ignore dep %s", fn)
		}
	}
	log.Debug("%s dep files %d", filename, count)
	return l, nil
}

func list(dst *Info, cmd, filename string) error {
	log.Debug("%s %s", cmd, filename)
	if out, err := utils.Exec(cmd, filename); err != nil {
		log.Debug("%s", err)
		return err
	} else {
		if err := parse(dst, string(out)); err != nil {
			return err
		}
	}
	return nil
}

func parse(dst *Info, src string) error {
	for _, line := range strings.Split(src, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "statically") {
			continue
		}
		//~ log.Print(line)
		items := strings.Split(line, " => ")
		ilen := len(items)
		tok0 := strings.TrimSpace(items[0])
		fn := ""
		if ilen == 1 {
			fn = strings.Split(tok0, " ")[0]
		} else {
			tok1 := strings.TrimSpace(items[1])
			fn = strings.Split(tok1, " ")[0]
		}
		_, ok := dst.Files[fn]
		if !ok {
			checkFile(dst, fn)
		}
	}
	return nil
}

func checkFile(dst *Info, fn string) {
	log.Debug("check file %s", fn)
	if filepath.IsAbs(fn) {
		dst.Files[fn] = true
	} else {
		dst.Files[fn] = false
	}
}
