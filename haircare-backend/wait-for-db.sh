#!/bin/sh
set -e

echo "Checking database readiness at $DB_HOST:$DB_PORT ..."

# Loop until Postgres accepts TCP connections
until nc -z "$DB_HOST" "$DB_PORT"; do
  echo "Database not ready, waiting..."
  sleep 2
done

echo "Database is ready — starting backend."
exec "$@"
