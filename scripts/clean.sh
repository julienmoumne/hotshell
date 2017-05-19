#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

go clean -i ./...

find . -type f -name 'bindata.go' -exec rm {} +

find ./cmd -type f -name 'mock_*.go' -exec rm {} +

./scripts/format.sh