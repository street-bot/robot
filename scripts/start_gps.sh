#!/bin/bash
set -eo pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR=$SCRIPT_DIR/..

echo "Starting GNSS service on EC25..."
python3 $ROOT_DIR/gps/start_gps_service.py
echo "GNSS service succesfully started on EC25"
rosrun nmea_navsat_driver nmea_serial_driver _port:=/dev/ttyUSB1 _baud:=115200