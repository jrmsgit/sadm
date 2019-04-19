// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package tpl

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"github.com/jrmsdev/sadm/internal/log"
)

func Parse(name, text string, data map[string]string) ([]byte, error) {
	log.Debug("%s", name)
	t, err := template.New(name).Parse(text)
	if err != nil {
		log.Debug("%s", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		log.Debug("%s", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

func parseFile(filename, name string, data map[string]string) ([]byte, error) {
	log.Debug("file %s", filename)
	blob, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Debug("%s", err)
		return nil, err
	}
	return Parse(name, string(blob), data)
}
