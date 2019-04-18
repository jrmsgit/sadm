// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jrmsdev/sadm/env"
	"github.com/jrmsdev/sadm/internal/cfg"
	"github.com/jrmsdev/sadm/internal/log"
)

var sprintf = fmt.Sprintf

var (
	cfgfile     string
	loglevel    string
	argsService string
)

var argsinit = map[string]string{
	"type":    "sadm",
	"service": "",
}

func init() {
	var (
		defCfg = filepath.Join(cfg.Prefix, "etc", "sadm.json")
		defLog = "error"
	)
	flag.StringVar(&cfgfile, "config", defCfg, "`file` path")
	flag.StringVar(&loglevel, "log", defLog, "`level`: debug, warn, error or quiet")
	// args init flags
	flag.StringVar(&argsService, "service", "", "service `name`")
}

var validCmd = map[string]bool{
	"pkg.list":    true,
	"fs.ls-mount": true,
}

func usage() {
	log.Print("usage: sadm-utils command [args...]")
	log.Print("commands:")
	for n := range validCmd {
		log.Printf("    %s", n)
	}
	log.Print("* run `sadm-utils -help` for more information")
}

func main() {
	flag.Parse()
	log.Init(loglevel)
	log.Debug("init")
	cmd := flag.Arg(0)
	var cmdargs []string
	if len(flag.Args()) < 1 {
		cmdargs = make([]string, 0)
	} else {
		cmdargs = flag.Args()[1:]
	}
	if cmd == "" {
		log.Errorf("no command specified")
		usage()
		os.Exit(1)
	} else if !validCmd[cmd] {
		log.Errorf("invalid command %s", cmd)
		usage()
		os.Exit(2)
	}
	dispatch(cmd, cmdargs)
}

func dispatch(cmd string, cmdargs []string) {
	log.Debug("dispatch %s %v", cmd, cmdargs)
	var (
		err    error
		config *cfg.Cfg
		opt    *env.Env
	)
	config, err = readConfig()
	if err != nil {
		log.Error(err)
		os.Exit(3)
	}
	argsinit["service"] = argsService
	opt, err = env.New(config, "sadm", argsinit)
	if err != nil {
		log.Error(err)
		os.Exit(4)
	}
	if cmd == "pkg.list" {
		err = pkgList(opt, cmdargs)
	} else if cmd == "fs.ls-mount" {
		err = fsLsMount(cmdargs)
	}
	if err != nil {
		log.Error(err)
		os.Exit(5)
	}
	os.Exit(0)
}

func readConfig() (*cfg.Cfg, error) {
	// read config file
	if fh, err := os.Open(cfgfile); err != nil {
		log.Error(err)
		return nil, err
	} else {
		// create config
		defer fh.Close()
		if config, err := cfg.New(fh); err != nil {
			log.Error(err)
			return nil, err
		} else {
			return config, nil
		}
	}
}
