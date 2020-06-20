#!/bin/bash
set -eo pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR=$SCRIPT_DIR/..

# Remove previously loaded v4l2loopback (if applicable)
sudo rmmod v4l2loopback

# Load v4l2loopback kernel module
sudo modprobe v4l2loopback video_nr=7,8

# Duplicate stream to other video devices
ffmpeg -i /dev/video0 -codec copy -f v4l2 /dev/video7 -codec copy -f v4l2 /dev/video8 &