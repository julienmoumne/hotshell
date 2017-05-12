#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

. ./scripts/common-vars.sh

go clean -i ./...

find . -type f -name 'bindata.go' -exec rm {} +

go fmt $ALL_BUT_VENDORS