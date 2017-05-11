#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

go get -u github.com/jteeuwen/go-bindata/...
go get -u github.com/kardianos/govendor
go get -u github.com/laher/goxc
sudo gem install github_changelog_generator