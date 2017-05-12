#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

ALL_BUT_VENDORS=$(go list ./... | grep -v /vendor/)