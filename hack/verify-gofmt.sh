#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

diff=$(gofmt -d -e ./ 2>&1) || true
if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Run 'make fmt'" >&2
  exit 1
fi
