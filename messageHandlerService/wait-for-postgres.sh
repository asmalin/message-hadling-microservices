#!/bin/bash
# wait-for-postgres.sh

set -e

host="$1"
user="$2"
db="$3"
shift 3
cmd="$@"

until pg_isready -h "$host" -U "$user" -d "$db"; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd