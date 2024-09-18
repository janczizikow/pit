#!/bin/bash
set -e

DB_DSN="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST/$DB_NAME?sslmode=disable"
migrate -path ./internal/database/migrations -database="$DB_DSN" up & PID=$!
wait $PID

/bin/api
