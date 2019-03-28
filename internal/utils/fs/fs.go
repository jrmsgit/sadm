// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"fmt"
	"os"

	"github.com/jrmsdev/sadm/internal/log"
)

var sprintf = fmt.Sprintf

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

func newFileInfo(filename string) (*Info, error) {
	s, err := os.Stat(filename)
	if err != nil {
		log.Debug("%s", err)
		return nil, err
	}
	return newInfo(s, filename), nil
}

func (i *Info) String() string {
	return i.filename
}

func (i *Info) Filename() string {
	return i.filename
}
