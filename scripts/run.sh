#!/bin/sh
set -e

/bin/api &
nginx -c "$PWD/nginx.conf" &
wait -n
