#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

./scripts/clean.sh
./scripts/install.sh