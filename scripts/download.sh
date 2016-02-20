#!/bin/bash
set -e

path="$1"
echo "Path used: $path"
name="$2"
echo "Name used: $name"
tracker="$3"
echo "Tracker used: $tracker"

cd $path
ctorrent -u '$tracker' '$name'
