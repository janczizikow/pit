#!/usr/bin/env bash
set -e

DB_DSN="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=$DB_SSL_MODE"
migrate -path ./internal/database/migrations -database="$DB_DSN" up &
wait -n

nginx -c "$PWD/nginx.conf" &
/usr/local/bin/app &

wait -n
