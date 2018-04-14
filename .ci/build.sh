#!/bin/bash

set -e -x

export GOPATH=$PWD/go

mkdir -p $GOPATH/src/go.qrs.fr
cp -r $1 $GOPATH/src/go.qrs.fr/$1

go test -v go.qrs.fr/$1

out_dir=$PWD/out_dir
mkdir -p $out_dir

for arch in amd64 386; do
	for os in windows linux darwin; do
		GOOS=$os GOARCH=$arch go build -o ${out_dir}/${1}_${os}_${arch} go.qrs.fr/$1
	done
done
