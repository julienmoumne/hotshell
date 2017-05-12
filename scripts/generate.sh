#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

go generate ./...