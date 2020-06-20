#!/bin/bash
set -eo pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR=$SCRIPT_DIR/..

cd $ROOT_DIR

source ./scripts/pull_binary.sh
source ./scripts/setup_devices.sh

docker-compose up -d