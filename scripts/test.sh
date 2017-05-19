#!/bin/bash
IFS=$'\n\t'
set -oeu pipefail

. ./scripts/common-vars.sh

./scripts/generate.sh

#todo find-out if '-cover' low percentages are due to bindata.go files
go test $ALL_BUT_VENDORS -timeout 10s