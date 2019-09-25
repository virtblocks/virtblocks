#!/bin/sh

echo "before (stdout)"
echo "before (stderr)" >&2

sleep 5

echo "after (stdout)"
echo "after (stderr)" >&2
