// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/nss"
)

func nssSync(ctx *env.Env, cmdargs []string) error {
	log.Debug("(%d) %v", len(cmdargs), cmdargs)
	if len(cmdargs) < 3 {
		return errors.New("invalid args")
	}
	destdir, err := filepath.Abs(cmdargs[0])
	if err != nil {
		log.Debug("%s", err)
		return err
	}
	ctx.Update("destdir", destdir)
	db := strings.TrimSpace(cmdargs[1])
	v := ""
	for _, k := range cmdargs[2:] {
		if k != "" {
			v = v + k + " "
		}
	}
	if err := ctx.Update("nss." + db, strings.TrimSpace(v)); err != nil {
		return err
	}
	return nss.Sync(ctx)
}
