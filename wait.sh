#!/bin/sh

set -e

host="go-exinity-task-postgress"
port="5432"

echo "⌛ Waiting for PostgreSQL at $host:$port..."

until pg_isready -h "$host" -p "$port"; do
  sleep 1
done

echo "✅ Postgres is ready — starting the app..."

exec "$@"
