// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package pkg

import (
	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
)

func Which(opt *env.Env, filename string) (*Info, error) {
	log.Debug("check %s", filename)
	var (
		m   Manager
		err error
	)
	m, err = newManager(opt)
	if err != nil {
		return nil, err
	}
	info := &Info{}
	err = m.Which(info, filename)
	if err != nil {
		return nil, err
	}
	log.Debug("which %s: %s", filename, info.Pkg)
	return info, nil
}
