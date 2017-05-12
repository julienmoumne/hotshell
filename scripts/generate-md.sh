#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

./scripts/install.sh

generate_md() {
    $GOPATH/bin/hs --generate-md --chdir -f $1 > $2
}

for e in examples/**/*.js; do
    generate_md $e $e.md
done

generate_md ./ COMMANDS.md