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

var prefix = "/usr/local"

type Cfg struct {
	EnvDir string
	CfgDir string
}

func New(src io.ReadCloser) (*Cfg, error) {
	log.Debug("new prefix=%s", prefix)
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
	config.CfgDir = filepath.Clean(config.CfgDir)
	//~ log.Debug("%#v", config)
	return config, nil
}

func setDefaults(config *Cfg) {
	p := filepath.FromSlash(prefix)
	config.EnvDir = filepath.Join(p, "etc", "sadm.d")
	config.CfgDir = filepath.Join(p, "etc", "sadm")
}
