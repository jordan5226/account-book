#!/bin/sh

set -e

echo "Run db migration"
/app/lib/pgdb/migration/main up

echo "Start app"
exec "$@"