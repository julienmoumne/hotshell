#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

read -r -p "dependency url: " url

go get $url
govendor add $url