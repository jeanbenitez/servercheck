export DB_HOST="localhost"
export DB_PORT=26257
export DB_NAME="servercheck"
export DB_USER="root"
export DB_PASS=""
export GOPORT=8005

# checking requered dependencies
if ! [ -x "$(command -v go)" ]; then
  echo 'Error: go is not installed.' >&2
  exit 1
fi
if ! [ -x "$(command -v dep)" ]; then
  echo 'Error: dep is not installed.' >&2
  exit 1
fi
if ! [ -x "$(command -v cockroach)" ]; then
  echo 'Error: cockroach is not installed.' >&2
  exit 1
fi
