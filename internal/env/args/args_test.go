// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package args

import (
	"testing"

	"github.com/jrmsdev/sadm/internal/_t/test"
)

var tsrc = map[string]string{
	"type": "sadm",
}

func TestMain(m *testing.M) {
	test.Main(m)
}

func TestArgs(t *testing.T) {
	cfg := test.NewConfig("{}")
	x, err := New(cfg, "test", tsrc)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(x)
}
