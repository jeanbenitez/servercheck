#! /bin/sh
set -e
source ./env.sh

# start cockroachdb
cockroach quit --insecure || true
cockroach start --insecure --listen-addr=$DB_HOST --background

# install and run
go install
servercheck
