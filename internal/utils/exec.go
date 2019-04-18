// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package utils

import (
	"context"
	"errors"
	"os/exec"
	"strings"
	"time"

	"github.com/jrmsdev/sadm/internal/log"
)

var ttl = 30 * time.Second

func Exec(command string, args ...string) ([]byte, error) {
	//~ log.Debug("%s %v", command, args)
	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()
	cmd := exec.CommandContext(ctx, command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Debug("%v %s", cmd.Args, err)
		return nil, errors.New(strings.TrimSpace(string(out)))
	}
	return out, nil
}
