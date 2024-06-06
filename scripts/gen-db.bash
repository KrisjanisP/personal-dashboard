#! /bin/bash

set -ex # Exit on error

SCRIPT_DIR=$(dirname $0)
pushd $SCRIPT_DIR
jet -source=sqlite -dsn="../data/sqlite3" -schema=dvds -path=./.gen