#!/bin/bash

set -e -x

dir=$1
commit=$2

DATE=`date +%Y-%m-%d`

for exe in `find ${dir} -type f`; do
	curl --ftp-create-dirs -T ${exe} ftp://skillz:YoloTest1234@skillz.events/${DATE}-${commit}/`basename ${exe}`
done
