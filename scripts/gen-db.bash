#! /bin/bash

set -ex # Exit on error

SCRIPT_DIR=$(dirname $0)
ROOT=$(readlink -f "$(dirname $SCRIPT_DIR)")
pushd $SCRIPT_DIR
jet -source=sqlite -dsn="$ROOT/data/sqlite3.db" -path="$ROOT/internal/database"
popd