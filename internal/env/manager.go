// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

type Manager interface {
	Dispatch(action string) (err error)
}
