#!/bin/sh
set -e

# Wait for PostgreSQL to be available
until nc -z postgres 5432; do
    echo "Waiting for PostgreSQL..."
    sleep 2
done

# Wait for Redis to be available
until nc -z redis 6379; do
    echo "Waiting for Redis..."
    sleep 2
done

# Run the main app
exec "$@"
