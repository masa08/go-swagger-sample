#!/bin/ash

set -e

host="$1"
shift
cmd="$@"

until psql "user=postgres password=postgres host=$host port=5432 dbname=lead" -c '\l'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd