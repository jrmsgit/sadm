// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"errors"

	"github.com/jrmsdev/sadm/internal/log"
	"github.com/jrmsdev/sadm/internal/utils/fs"
)

func fsLsMount(args []string) error {
	if len(args) < 1 {
		return errors.New("destdir not set")
	}
	d := args[0]
	ls, err := fs.LsMount(d)
	if err != nil {
		return err
	}
	for _, mnt := range ls {
		log.Print(mnt)
	}
	return nil
}
