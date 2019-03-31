// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

import (
	fp "path/filepath"
	"testing"

	"github.com/jrmsdev/sadm/internal/_t/test"
)

func TestMain(m *testing.M) {
	test.Main(m)
}

func TestCfg(t *testing.T) {
	x, err := New(test.NewReadCloser("{}"))
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
