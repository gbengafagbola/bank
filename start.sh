#!/bin/sh --> file would be run by /bin/sh since it an alpine image and bash isn't available

set -e 

echo "run db migration"  
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"