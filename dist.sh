#!/bin/sh -eu
BUILD_NUMBER=${APPVEYOR_BUILD_NUMBER:-''}
PREFIX=${1:-'prefix'}
pwd
mkdir -p build
cd ${PREFIX} && {
	tar -vcJf ../build/sadm-build${BUILD_NUMBER}.txz .
}
cd ../build
sha512sum sadm-build*.txz >sadm-build${BUILD_NUMBER}-sha512sum.txt
