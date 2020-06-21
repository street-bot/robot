# 4G/LTE Connection Issue Cause Analysis

## Symptoms
- SSH session freezes (then gets closed by remote host) on large input/output (partially solved by 1.)
- WebSocket intermittently opens and closes causing WebRTC handshake to fail
- ~10s after initiation of a successful WebRTC connection, the remote host disconnects from both SSH and WebRTC, suggesting that the LTE may have disconnected.
- LTE disconnects when upload speed is beyond a certain threshold



## Actions with negative results
- Increase sleep time on `lte_connect` script from 30s -> 600s. Did not make any noticable difference.
- Switched USB-C cable between NUC and LTE module. Didn't make any noticable difference.
- Added `AlwaysExitOnFailure` on the `streetbot_ssh` service. Didn't any noticable difference.
- Checked connection ECIO, RSSI, SINR and they are all within "Excellent" ranges.
- Ping to `google.ca` doesn't show any packet loss.
- `qmicli -d /dev/cdc-wdm0 --wds-get-packet-statistics` does not show any dropped packets on TRX.
- Changing the MTU causes Docker TLS handshakes to timeout. Making the Docker MTU the same as the WWAN0's MTU restores TLS handshake capability.

## Actions with positive results
1. Changing SSH relay to a native machine on DO Droplet improves SSH reconnection speed.
2. Decreased MTU from 1500 -> 1000 seemed to improve the SSH package connection robustness with regards to streaming large packets.
3. Turning on MSS clamping restores Docker TLS handshake capability after lowering WWAN0 MTU.
4. Reduced TX power via AT commands (23 dbm -> 22 dbm). Video was working without dropping connection.
5. Changed MTU setting on `docker-compose` file to 1300. 

## Findings
- MTU settings will affect docker bridge TLS connections. 

### AT Commands
```bash
AT+QNVFW="/nv/item_files/rfnv/00020993",0100D200
AT+QNVFW="/nv/item_files/rfnv/00020995",0100D200
AT+QNVFW="/nv/item_files/rfnv/00022191",0100D200

AT+QNVFR="/nv/item_files/rfnv/00020993"
AT+QNVFR="/nv/item_files/rfnv/00020995"
AT+QNVFR="/nv/item_files/rfnv/00022191"
```