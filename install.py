#!/usr/bin/env python3

import os
import sys

from argparse import ArgumentParser

prefix = './prefix'

args = ArgumentParser(description = 'sadm installer')
args.add_argument('--prefix', default = prefix, metavar = prefix,
	help = 'install prefix')
args.add_argument('--remove', action = 'store_true', default = False,
	help = 'uninstall')

def call(cmd):
	rc = os.system(cmd)
	if rc != 0:
		print(cmd, 'failed')
		sys.exit(rc)

def build_cfg(src_file):
	dst_file = src_file[:-3]
	with open(src_file, 'r') as src:
		with open(dst_file, 'w') as dst:
			for line in src.readlines():
				line = line.replace('[[PREFIX]]', prefix)
				dst.write(line)
	print('%s done' % dst_file)

def install():
	os.environ['GOBIN'] = '%s/bin' % prefix

	call('rm -f ./internal/log/debug.go')
	call('cd ./lib && ./gen.sh --prefix %s/lib/sadm' % prefix)

	build_cfg('./internal/cfg/build.go.in')
	build_cfg('./etc/sadm.json.in')

	call('go install ./bin/sadm')
	print('%s/bin/sadm installed' % prefix)

	call('mkdir -p %s/etc/sadm' % prefix)
	call('cp -va etc/sadm.json %s/etc' % prefix)

	call('mkdir -p %s/share/doc/sadm/examples' % prefix)
	call('cp -va etc/sadm/*.json %s/share/doc/sadm/examples' % prefix)
	call('cp -va etc/sadm/template %s/share/doc/sadm/examples' % prefix)

def uninstall():
	cmd = 'rm -rfv %s/bin/sadm' % prefix
	cmd += ' %s/etc/sadm' % prefix
	cmd += ' %s/etc/sadm.json' % prefix
	cmd += ' %s/share/doc/sadm' % prefix
	call(cmd)

if __name__ == '__main__':
	flags = args.parse_args()
	prefix = os.path.abspath(flags.prefix)
	if flags.remove:
		uninstall()
	else:
		install()
	sys.exit(0)
