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

var Prefix = "/usr/local"

type Cfg struct {
	LibDir string
	CfgDir string
}

func New(src io.ReadCloser) (*Cfg, error) {
	log.Debug("new")
	config := new(Cfg)
	setDefaults(config)
	log.Debug("prefix=%s", Prefix)
	defer src.Close()
	if blob, err := ioutil.ReadAll(src); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(blob, &config); err != nil {
			return nil, err
		}
	}
	// paths cleanup
	config.LibDir = filepath.Clean(config.LibDir)
	config.CfgDir = filepath.Clean(config.CfgDir)
	//~ log.Debug("%#v", config)
	return config, nil
}

func setDefaults(config *Cfg) {
	config.LibDir = filepath.Join(Prefix, "lib", "sadm")
	config.CfgDir = filepath.Join(Prefix, "etc", "sadm")
}
