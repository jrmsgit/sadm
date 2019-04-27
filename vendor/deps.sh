#!/bin/sh -eu
deps='github.com/jrmsdev/gojc'
for repo in $(echo ${deps}); do
	echo ${repo}
	mkdir -p ./vendor/${repo}
	git clone --depth 3 https://${repo}.git ./vendor/${repo}
done
