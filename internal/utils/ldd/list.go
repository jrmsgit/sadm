// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ldd

import (
	"context"
	"os/exec"
	"strings"
	"time"

	"github.com/jrmsdev/sadm/internal/env/args"
	"github.com/jrmsdev/sadm/internal/log"
)

func List(opt *args.Args, filename string) ([]string, error) {
	x := opt.Get("ldd")
	if x == "" {
		x = "ldd"
	}
	log.Debug("%s %s", x, filename)
	l := make([]string, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, x, filename)
	if out, err := cmd.CombinedOutput(); err != nil {
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
