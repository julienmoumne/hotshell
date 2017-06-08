#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

. ./scripts/common-vars.sh

go vet $ALL_BUT_VENDORS