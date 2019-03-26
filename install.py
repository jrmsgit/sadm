#!/usr/bin/env python3

import os
import sys

from argparse import ArgumentParser

prefix = '/usr/local'

def build_cfg(src_file):
	dst_file = src_file[:-3]
	with open(src_file, 'r') as src:
		with open(dst_file, 'w') as dst:
			for line in src.readlines():
				line = line.replace('[[PREFIX]]', prefix)
				dst.write(line)
	print('%s done' % dst_file)

args = ArgumentParser(description = 'sadm installer')
args.add_argument('--prefix', default = prefix, metavar = prefix, help = 'install prefix')
args.add_argument('--remove', action = 'store_true', default = False, help = 'uninstall')

def install():
	os.environ['GOBIN'] = '%s/bin' % prefix

	build_cfg('./internal/cfg/build.go.in')
	build_cfg('./etc/sadm.json.in')

	os.system('go install ./bin/sadm')
	print('%s/bin/sadm installed' % prefix)

	os.system('mkdir -vp %s/etc/sadm.d' % prefix)
	os.system('cp -va etc/sadm.json %s/etc' % prefix)

	os.system('mkdir -vp %s/etc/sadm' % prefix)
	os.system('cp -va etc/sadm/* %s/etc/sadm' % prefix)

def uninstall():
	os.system('rm -rfv %s/bin/sadm %s/etc/sadm.d %s/etc/sadm %s/etc/sadm.json' % (prefix, prefix, prefix, prefix))

if __name__ == '__main__':
	flags = args.parse_args()
	prefix = flags.prefix
	if flags.remove:
		uninstall()
	else:
		install()
	sys.exit(0)
