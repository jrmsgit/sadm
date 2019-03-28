// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"errors"
	"os"
)

var DiffErr = errors.New("diff error")

type Info struct {
	os.FileInfo
	filename string
}

func newInfo(base os.FileInfo, filename string) *Info {
	return &Info{
		base,
		filename,
	}
}

func (i *Info) String() string {
	return i.filename
}

func (i *Info) Filename() string {
	return i.filename
}
