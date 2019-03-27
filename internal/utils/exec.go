// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package utils

import (
	"context"
	"os/exec"
	"time"

	"github.com/jrmsdev/sadm/internal/log"
)

var ttl = 100 * time.Millisecond

func Exec(command string, args ...string) ([]byte, error) {
	log.Debug("%s %v", command, args)
	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()
	cmd := exec.CommandContext(ctx, command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Debug("%s", err)
		return nil, err
	}
	return out, nil
}
