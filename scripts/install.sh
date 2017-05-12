#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

. ./scripts/common-vars.sh

./scripts/generate.sh

go vet $ALL_BUT_VENDORS

go install ./...

which hs
$GOPATH/bin/hs -v