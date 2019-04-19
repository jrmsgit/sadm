// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package tpl

import (
	"bytes"
	"text/template"

	"github.com/jrmsdev/sadm/internal/log"
)

func Parse(text string, data map[string]string) (string, error) {
	log.Debug("%s", text)
	t, err := template.New("sadm").Parse(text)
	if err != nil {
		log.Debug("%s", err)
		return "", err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		log.Debug("%s", err)
		return "", err
	}
	return buf.String(), nil
}
