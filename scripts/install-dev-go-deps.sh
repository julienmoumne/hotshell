#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

go get -u github.com/jteeuwen/go-bindata/...
go get -u github.com/kardianos/govendor
go get -u github.com/laher/goxc
go get -u github.com/vektra/mockery/...