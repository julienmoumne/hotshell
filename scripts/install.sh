#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

. ./scripts/common-vars.sh

./scripts/generate.sh

./scripts/vet.sh

go install ./...

which hs
$GOPATH/bin/hs -v