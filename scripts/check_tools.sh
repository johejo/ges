#!/usr/bin/env bash

set -euC pipefail

TOOLS="go golint goimports mockgen wire docker openapi-generator"

for cmd in ${TOOLS}; do
  printf "%-20s" "$cmd"
  if type "$cmd" >/dev/null 2>&1; then
    echo OK
  else
    echo NG
  fi
done
