#!/bin/sh

test "$#" = 2 || exit 1

echo "$1 (stdout)"
echo "$1 (stderr)" >&2

sleep 5

echo "$2 (stdout)"
echo "$2 (stderr)" >&2
