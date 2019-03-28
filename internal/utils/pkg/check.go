// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func Check(opt *args.Args, filename string) error {
	log.Debug("check %s", filename)
	var (
		m       Manager
		err     error
	)
	m, err = newManager(opt)
	if err != nil {
		return err
	}
	info := &Info{}
	err = m.Which(info, filename)
	if err != nil {
		return err
	}
	log.Debug("%s provided by %s", filename, info.Pkg)
	err = m.Depends(info)
	if err != nil {
		return err
	}
	log.Debug("%s depends on %v", info.Pkg, info.Deps)
	return nil
}
