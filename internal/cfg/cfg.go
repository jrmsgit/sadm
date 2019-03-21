// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/log"
)

type Cfg struct {
	EnvDir string
}

func New(src io.ReadCloser) (*Cfg, error) {
	log.Debug("new")
	config := new(Cfg)
	config.EnvDir = filepath.FromSlash("/usr/local/etc/sadm.d")
	defer src.Close()
	if blob, err := ioutil.ReadAll(src); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(blob, &config); err != nil {
			return nil, err
		}
	}
	config.EnvDir = filepath.Clean(config.EnvDir)
	log.Debug("%s", config)
	return config, nil
}

func (c *Cfg) String() string {
	return fmt.Sprintf("%#v", c)
}
