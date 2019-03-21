#!/usr/bin/env python3

import os
import sys

prefix = '/usr/local'

if __name__ == '__main__':
	os.environ['GOBIN'] = '%s/bin' % prefix

	os.system('go install ./bin/sadm')
	print('%s/bin/sadm installed' % prefix)

	os.system('mkdir -vp %s/etc/sadm.d' % prefix)
	os.system('cp -va etc/sadm.json %s/etc' % prefix)

	os.system('mkdir -vp %s/lib/sadm' % prefix)
	os.system('cp -va lib %s/lib/sadm' % prefix)

	sys.exit(0)
