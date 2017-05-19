#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

./scripts/install-dev-go-deps.sh

sudo gem install github_changelog_generator