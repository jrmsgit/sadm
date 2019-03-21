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
	name string
	args map[string]string
}

func New(name string, src io.ReadCloser) (*Env, error) {
	log.Debug("new %s", name)
	environ := new(Env)
	environ.name = name
	defer src.Close()
	if blob, err := ioutil.ReadAll(src); err != nil {
		return nil, err
	} else {
		environ.args = make(map[string]string)
		if err := json.Unmarshal(blob, &environ.args); err != nil {
			return nil, err
		}
	}
	log.Debug("%#v", environ)
	return environ, nil
}

func (e *Env) String() string {
	return e.name
}

func (e *Env) Type() string {
	return e.args["type"]
}
