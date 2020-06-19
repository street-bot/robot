#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR=$SCRIPT_DIR/..

pushd /tmp

# Pull release binary from GitHub
GITHUB_TOKEN=$(cat /home/streetbot/.github_token)
OWNER="street-bot"
REPO="robot"
RELEASE_ID=$(curl -H "Authorization: token $GITHUB_TOKEN" -sL https://api.github.com/repos/street-bot/robot/releases/latest | jq -r ".assets[] | select(.name | contains(\"robot_linux_amd64.tar.gz\")) | .id")
rm -f robot_linux_amd64.tar.gz
curl -H 'Accept: application/octet-stream' -H "Authorization: token $GITHUB_TOKEN" -LJO "https://api.github.com/repos/$OWNER/$REPO/releases/assets/$RELEASE_ID"

# Decompress and copy binary to robot directory
tar -xvf ./robot_linux_amd64.tar.gz
chmod +x ./robot
mv ./robot $ROOT_DIR/robot

popd