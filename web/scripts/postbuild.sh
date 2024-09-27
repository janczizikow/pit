#!/bin/bash
set -e

# remove all source maps
find "$(pwd)"/build -name "*.js.map" -type f -delete

# move 404 page from 404 folder to root
mv "$(pwd)/build/404/index.html" "$(pwd)/build/404.html"
rm -rf "$(pwd)/build/404"

sed -i -e 's/\.\.\//.\//g' build/404.html
rm "$(pwd)/build/404.html-e"
