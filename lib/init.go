// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package lib

import (
	"archive/zip"
	"bytes"
	"encoding/base64"

	"github.com/jrmsdev/sadm/internal/log"
)

var prefix string
var zipfile string
var storage map[string]*zip.File

var b64 = base64.StdEncoding.DecodeString

func Init() error {
	if prefix == "" {
		prefix = "./lib"
	}
	log.Debug("prefix %s", prefix)
	log.Debug("lib.zip.b64 %d", len(zipfile))
	storage = make(map[string]*zip.File)
	if zipfile != "" {
		blob, err := b64(zipfile)
		if err != nil {
			log.Debug("%s", err)
			return err
		}
		zdata := bytes.NewReader(blob)
		var zr *zip.Reader
		log.Debug("lib.zip %d", zdata.Len())
		zr, err = zip.NewReader(zdata, int64(zdata.Len()))
		if err != nil {
			log.Debug("%s", err)
			return err
		}
		for _, f := range zr.File {
			storage[f.Name] = f
			log.Debug("load %s", f.Name)
		}
	}
	return nil
}
