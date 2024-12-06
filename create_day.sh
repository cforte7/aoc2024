#!/bin/bash

# Check if a number is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <number>"
  exit 1
fi

n="$1"

# Check if templ.go exists
if [ ! -f "templ.go" ]; then
  echo "Error: templ.go not found."
  exit 1
fi

# Create the folder
folder="day$n"
mkdir -p "$folder"

# Create the files
cp "templ.go" "$folder/day$n.go"
touch "$folder/test.txt" "$folder/input.txt"

echo "Folder '$folder' and files created successfully."
