// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"os"
	"testing"

	"github.com/jrmsdev/sadm/internal/log"
)

func Main(m *testing.M) {
	if err := log.Init("quiet"); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
