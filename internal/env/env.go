// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/jrmsdev/sadm/internal/log"
)

type Env struct {
	Name string `json:"-"`
}

func New(name string, src io.ReadCloser) (*Env, error) {
	log.Debug("new %s", name)
	environ := new(Env)
	environ.Name = name
	defer src.Close()
	if blob, err := ioutil.ReadAll(src); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(blob, &environ); err != nil {
			return nil, err
		}
	}
	log.Debug("%#v", environ)
	return environ, nil
}

func (e *Env) String() string {
	return e.Name
}

func (e *Env) Run(action string) error {
	log.Debug("run %s %s", action, e)
	return nil
}
