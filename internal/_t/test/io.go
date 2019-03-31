// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"bytes"
)

type ReadCloser struct {
	*bytes.Buffer
}

func NewReadCloser(s string) *ReadCloser {
	return &ReadCloser{bytes.NewBufferString(s)}
}

func (r *ReadCloser) Close() error {
	return nil
}
