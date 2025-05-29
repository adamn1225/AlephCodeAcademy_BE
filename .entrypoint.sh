#!/bin/sh
until pg_isready -h db -p 5432 -U aleph; do
  echo "Waiting for postgres..."
  sleep 1
done

exec "$@"
