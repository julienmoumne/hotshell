#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

github_changelog_generator

./scripts/generate-md.sh