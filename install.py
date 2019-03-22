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
args.add_argument('--prefix', default = prefix)

if __name__ == '__main__':
	flags = args.parse_args()
	prefix = flags.prefix

	os.environ['GOBIN'] = '%s/bin' % prefix

	build_cfg('./internal/cfg/build.go.in')
	build_cfg('./etc/sadm.json.in')

	os.system('go install ./bin/sadm')
	print('%s/bin/sadm installed' % prefix)

	os.system('mkdir -vp %s/etc/sadm.d' % prefix)
	os.system('cp -va etc/sadm.json %s/etc' % prefix)

	os.system('mkdir -vp %s/lib/sadm' % prefix)
	os.system('cp -va lib %s/lib/sadm' % prefix)

	sys.exit(0)
