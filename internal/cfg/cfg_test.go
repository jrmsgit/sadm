// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

import (
	"bytes"
	fp "path/filepath"
	"os"
	"testing"

	"github.com/jrmsdev/sadm/internal/log"
)

type readCloser struct {
	*bytes.Buffer
}

func newReadCloser(s string) *readCloser {
	return &readCloser{bytes.NewBufferString(s)}
}

func (r *readCloser) Close() error {
	return nil
}

func TestMain(m *testing.M) {
	if err := log.Init("quiet"); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestCfg(t *testing.T) {
	x, err := New(newReadCloser("{}"))
	if err != nil {
		t.Fatal(err)
	}
	if x.LibDir != fp.Join(Prefix, "lib", "sadm") {
		t.Errorf("LibDir %s", x.LibDir)
	}
	if x.CfgDir != fp.Join(Prefix, "etc", "sadm") {
		t.Errorf("CfgDir %s", x.CfgDir)
	}
}
