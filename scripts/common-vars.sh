#!/bin/bash
IFS=$'\n\t'
set -oeu pipefail

ALL_BUT_VENDORS=$(go list ./... | grep -v /vendor/)