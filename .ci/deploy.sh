#!/bin/bash

set -e -x

dir=$1
commit=$2

DATE=`date +%Y-%m-%d`

for exe in ${dir}; do
	curl --ftp-create-dirs -T ${dir}/${exe} ftp://skillz:YoloTest1234@skillz.events/${DATE}-${commit}/${exe}
done
