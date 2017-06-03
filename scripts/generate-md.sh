#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

generate_md() {
    hsFile=$(basename $1)
    mdFile=$(basename $2)
    (cd $(dirname $1); $GOPATH/bin/hs --generate-md -f $hsFile > $mdFile)
}

for e in examples/**/*.js; do
    generate_md $e $e.md
done

generate_md ./ COMMANDS.md