#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

$GOPATH/bin/hs-man | gzip > debian/usr/share/man/man1/hs.1.gz