// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/jrmsdev/sadm/internal/cfg"
	"github.com/jrmsdev/sadm/internal/env"
	"github.com/jrmsdev/sadm/internal/log"
)

var (
	cfgfile  string
	loglevel string
)

func init() {
	var (
		defCfg = filepath.Join(cfg.Prefix, "etc", "sadm.json")
		defLog = "error"
	)
	flag.StringVar(&cfgfile, "config", defCfg, "`file` path")
	flag.StringVar(&loglevel, "log", defLog, "`level`: debug, warn, error or quiet")
}

func usage() {
	log.Print("usage: sadm env action")
	log.Print("*run `sadm -help` for more information")
}

func main() {
	// parse args and init log
	flag.Parse()
	log.Init(loglevel)
	log.Debug("init")
	envname := flag.Arg(0)
	action := flag.Arg(1)
	argsok := true
	if envname == "" {
		log.Errorf("no env name")
		argsok = false
	} else if action == "" {
		log.Errorf("no action")
		argsok = false
	} else if action != "dump" && action != "check" {
		log.Errorf("invalid action %s", action)
		argsok = false
	}
	if !argsok {
		usage()
		os.Exit(9)
	}
	// read config file
	if fh, err := os.Open(cfgfile); err != nil {
		log.Error(err)
		os.Exit(1)
	} else {
		// create config
		defer fh.Close()
		if config, err := cfg.New(fh); err != nil {
			log.Error(err)
			os.Exit(2)
		} else {
			// load env
			envfile := filepath.Join(config.EnvDir, envname+".json")
			if envfh, err := os.Open(envfile); err != nil {
				log.Error(err)
				os.Exit(5)
			} else {
				// create env
				if environ, err := env.New(config, envname, envfh); err != nil {
					log.Error(err)
					os.Exit(3)
				} else {
					// run env action
					if err := env.Run(environ, action); err != nil {
						log.Error(err)
						os.Exit(4)
					}
				}
			}
		}
	}
	os.Exit(0)
}
