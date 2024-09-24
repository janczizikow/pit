#!/usr/bin/env bash
set -e

# /bin/api
nginx -c "$PWD/nginx.conf" &
/bin/api &

wait -n
