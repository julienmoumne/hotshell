#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

. ./scripts/common-vars.sh

go fmt $ALL_BUT_VENDORS