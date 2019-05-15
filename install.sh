#! /bin/sh
set -e

source ./env.sh

# Install dependencies (dep required)

dep ensure -v -update

# Start and load database schema (cockroachdb required)

cockroach quit --insecure || true

cockroach start --insecure --listen-addr=$DB_HOST --background

cockroach sql --insecure \
--execute="CREATE DATABASE IF NOT EXISTS $DB_NAME" \

cockroach sql --insecure \
--database=$DB_NAME < $GOPATH/src/github.com/jeanbenitez/servercheck/resources/schema.sql \

cockroach quit --insecure
