#!/bin/bash
IFS=$'\n\t'
set -oeu pipefail

. ./scripts/common-vars.sh

go generate $ALL_BUT_VENDORS