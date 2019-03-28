// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func Check(opt *args.Args, filename string) (*Info, error) {
	log.Debug("check %s", filename)
	var (
		err  error
		info *Info
	)
	info, err = Which(opt, filename)
	if err != nil {
		return nil, err
	}
	err = List(opt, info)
	if err != nil {
		return nil, err
	}
	err = fs.Check(opt, info.Files...)
	if err != nil {
		return nil, err
	}
	return info, nil
}
