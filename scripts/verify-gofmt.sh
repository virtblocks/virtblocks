#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

diff=$(gofmt -l ./ 2>&1 | grep -Ev '^_obj/') || true
if [[ -n "${diff}" ]]; then
  echo "The following files are not formatted properly:" >&2
  echo >&2
  echo "${diff}" >&2
  echo >&2
  echo "Run 'make fmt' to fix them." >&2
  exit 1
fi
