#!/bin/bash

set -e -x

#go get -u github.com/go-bindata/go-bindata/...
#cd $GOPATH/src/go.qrs.fr/$1
#go-bindata ./static
#cd -

go test -v go.qrs.fr/$1

out_dir=$2
mkdir -p $out_dir

for arch in amd64 386; do
	for os in linux darwin; do
		GOOS=$os GOARCH=$arch go build -ldflags="-s" -o ${out_dir}/${1}_${os}_${arch} go.qrs.fr/$1
	done
done

for arch in amd64 386; do
	GOOS=windows GOARCH=$arch go build -ldflags="-s" -o ${out_dir}/${1}_windows_${arch}.exe go.qrs.fr/$1
done
