#!/bin/bash
IFS=$'\n\t'
set -oeu pipefail

VERSION=$(cat VERSION)

read -r -p "is $VERSION the new release? [y/N] " response
if ! [[ "$response" =~ ^([yY][eE][sS]|[yY])+$ ]]
then
    exit 0
fi

set -x
./scripts/clean-install.sh
./scripts/generate-man.sh
goxc -pv $VERSION -wd ./cmd/hs