#!/usr/bin/env python3

import os
import sys

if __name__ == '__main__':
	os.environ['GOBIN'] = '/opt/sadm/bin'
	os.system('go install ./bin/sadm')
	print('/opt/sadm/bin/sadm installed')
	os.system('mkdir -vp /opt/sadm/lib/sadm')
	os.system('cp -va lib /opt/sadm/lib/sadm')
	sys.exit(0)
