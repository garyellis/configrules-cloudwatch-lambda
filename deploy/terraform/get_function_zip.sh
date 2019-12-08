#!/bin/bash

eval "$(jq -r '@sh "URL=\(.url) OUT=\(.out)"')"

curl -s -L -o $OUT $URL

jq -n --arg path "$OUT" '{"path":$path}'
