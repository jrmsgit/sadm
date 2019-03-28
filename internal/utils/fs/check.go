// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func Check(opt *args.Args, files ...string) error {
	log.Debug("check %s: %d files", opt.Env, len(files))
	return nil
}
