#!/bin/bash
set -eo pipefail

# Current script directory
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR=$SCRIPT_DIR/..

pushd $ROOT_DIR

docker build . -t registry.digitalocean.com/streetbot/ros:latest
docker push registry.digitalocean.com/streetbot/ros:latest

popd
