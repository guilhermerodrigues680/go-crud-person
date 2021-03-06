#!/bin/bash

# The main file
CMD_MAIN="cmd/main.go"

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
DIR="$( cd "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

# Delete the old bin dir
echo "==> Removing old bin directory..."
rm -rf bin/
mkdir -p bin/

echo "==> Building ..."
echo "64-Bit - Linux"
GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 $CMD_MAIN

echo
echo "==> Done!"
echo "==> Results:"
ls -lh bin/
echo
