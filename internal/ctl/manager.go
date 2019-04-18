// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ctl

type Manager interface {
	Dispatch(action string) (err error)
}
