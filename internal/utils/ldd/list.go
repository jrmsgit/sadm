// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ldd

import (
	"strings"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils"
)

func List(opt *args.Args, filename string) ([]string, error) {
	x := opt.Get("ldd")
	if x == "" {
		x = "ldd"
	}
	log.Debug("%s %s", x, filename)
	l := make([]string, 0)
	if out, err := utils.Exec(x, filename); err != nil {
		log.Debug("%s", err)
		return nil, err
	} else {
		if err := parse(l, string(out)); err != nil {
			return nil, err
		}
	}
	return l, nil
}

func parse(dst []string, src string) error {
	log.Debug("parse")
	for _, line := range strings.Split(src, "\n") {
		line = strings.TrimSpace(line)
		log.Print(line)
	}
	return nil
}
