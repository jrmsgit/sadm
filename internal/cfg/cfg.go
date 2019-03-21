// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/log"
)

type Cfg struct {
	EnvDir string
	LibDir string
}

func New(src io.ReadCloser) (*Cfg, error) {
	log.Debug("new")
	config := new(Cfg)
	setDefaults(config)
	defer src.Close()
	if blob, err := ioutil.ReadAll(src); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(blob, &config); err != nil {
			return nil, err
		}
	}
	// paths cleanup
	config.EnvDir = filepath.Clean(config.EnvDir)
	config.LibDir = filepath.Clean(config.LibDir)
	//~ log.Debug("%#v", config)
	return config, nil
}

func setDefaults(config *Cfg) {
	config.EnvDir = filepath.FromSlash("/usr/local/etc/sadm.d")
	config.LibDir = filepath.FromSlash("/usr/local/lib/sadm")
}
