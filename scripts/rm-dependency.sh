#!/bin/bash
IFS=$'\n\t'
set -oxeu pipefail

read -r -p "dependency url: " url

govendor remove $url