#!/bin/bash

set -e -u -x

export GOPATH=$PWD/gopath
export PATH=$PWD/gopath/bin:$PATH

cd cf-demos/blue-green

echo
echo "Running build..."
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ../../artefacts/app . 
cp ci/manifest.yml ../../artefacts/
