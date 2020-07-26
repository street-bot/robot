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
5. Changed MTU setting on `docker-compose` file to 1300. MTU was later increased ot 1400 with no observed changes.

## Findings
- MTU settings will affect docker bridge TLS connections.
- The current USB LTE module power deliver may be insufficient causing module brown/black-out when higher broadcasting settings are needed.
- Signal strength does not seem to be an issue. The RSSI, ECIO, and SINRs have been consistent; we've also moved the module rigth up to the window.

### AT Commands
```bash
AT+QNVFW="/nv/item_files/rfnv/00020993",0100D200
AT+QNVFW="/nv/item_files/rfnv/00020995",0100D200
AT+QNVFW="/nv/item_files/rfnv/00022191",0100D200

AT+QNVFR="/nv/item_files/rfnv/00020993"
AT+QNVFR="/nv/item_files/rfnv/00020995"
AT+QNVFR="/nv/item_files/rfnv/00022191"
```

## Actions with no noticable impact:
### Enable Debug Mode on the Intel Wirelss Card -> reverted
```
    echo 1 > /sys/kernel/debug/iwlwifi/0000\:0X\:00.0/iwlmvm/fw_dbg_collect

    cat /sys/devices/virtual/devcoredump/devcdY/data > iwl.dump
    echo 1 > /sys/devices/virtual/devcoredump/devcdY/data

    cat << EOF > /etc/udev/rules.d/85-iwl-dump.rules
    SUBSYSTEM=="devcoredump", ACTION=="add", RUN+="/usr/local/sbin/iwlfwdump.sh"
    EOF

    cat << EOF > /usr/local/sbin/iwlwdump.sh
    #!/bin/bash
    timestamp=$(date +%F_%H-%M-%S)
    filename=/var/log/iwl-fw-error_${timestamp}.dump
    cat /sys/${DEVPATH}/data > ${filename}
    echo 1 > /sys/${DEVPATH}/data
    EOF
    chmod a+x /usr/local/sbin/iwlwdump.sh
    Now when firmware crashes I should see something like this:

        iwlwifi 0000:01:00.0: Microcode SW error detected.  Restarting 0x82000000.
        [snip]
        iwlwifi 0000:01:00.0: Loaded firmware version: XX.XX.XX.XX
        iwlwifi 0000:01:00.0: 0x0000090A | ADVANCED_SYSASSERT
```

### Disable Power Save on the WIFI Nic -> reverted
    echo "options iwlmvm power_scheme=1" >> /etc/modprobe.d/iwlwifi.conf

## Additional Actions with Positive Impact

### Problem: intel wifi driver errors in syslog and very slow boottime
### Fix: Upgraded to an HWE Kernel
    This Fixed the first intel wifi error in dmesg

### Problem: unandled alg 0x007 in dmesg related to intel wifi
### Fix: Try an older intel wifi firmware with a backported Fix -> this made the errors in syslog go away
    sudo apt install git build-essential
    git clone https://git.kernel.org/pub/scm/linux/kernel/git/iwlwifi/backport-iwlwifi.git
    cd backport-iwlwifi
    make defconfig-iwlwifi-public
    sed -i 's/CPTCFG_IWLMVM_VENDOR_CMDS=y/# CPTCFG_IWLMVM_VENDOR_CMDS is not set/' .config
    make -j4
    sudo make install
    cd /lib/firmware


### Problem:
        1. ssh -J ssh-relay.street-bot.com -p 2223 streetbot@localhost
        2. speedtest-cli
        3. SSH sesssion dies
### Fix:
        1. /etc/ssh/sshd_config  -> This fixed the problem when the NUC was connected to WIFI
            -> Changed ClientAliveInterval from 1 to 600
            -> Added ClientAliveCountMax to 10
            -> And Reboot

        2. Set TCPKeepAlive to no in /etc/ssh/sshd_config -> reverted. This didn't make a difference

        3. Refactored the autossh server -> see /etc/systemd/system/streetbot_ssh
                                         -> see /etc/default/autossh


        4. On the SSH Relay: -> This let the speed test complete consistently over LTE

            In /etc/sshd_config increased the ClientAliveInterval from 1 to 60
            and restart ssh

### Problem: Unpredictable Network Interface Names
### Fix:
        1. Add  "net.ifnames=0 biosdevname=0" to to GRUB_CMDLIN_LINX="" line
           in /etc/default/grub
        2. sudo grub-mkconfig -o /boot/grub/grub.cfg


### Now that the speedtest runs over the LTE let's improve the upload Speed

### Crank up the TX Transmit Power on the Quectel Module

    1. Fido uses LTE Bands 4, 7 and 17
        4:      00020995
        7:      00020993
        17:     00023133

        - Quectel Module Only Supports Band 4 out of these options


    2. To Read TX Power Settings:

        AT+QNVFR="/nv/item_files/rfnv/<band setting id>"
        AT+QNVFR="/nv/item_files/rfnv/00020995      -> Set to 0100D200 (got us 3 Mbit/s upload)


    3. To Write TX Power Settings:

        AT+QNVFW="/nv/item_files/rfnv/<band setting id>",<power setting>
        AT+QNVFW="/nv/item_files/rfnv/00020995",0100F000

    4. Result

        Didn't see a significant increase in performance (3.2 Mbit/s compared to 2.9 Mbit/s before change)
        But ! The LTE where I was testing wasn't amazing:
            RSSI: -82 dBm
            ECIO: -2.5 dBm
            SINR: 9.0 dB
            RSRQ: -8 dBm
            RSRP: -105 dBm

### Problem: Docker renaming network interfaces causes NIC bounce
### Possible Fix: (Reverted. It didn't work)
```
cat << EOF > /etc/systemd/network/99-default.link
[Match]
Path=/devices/virtual/net/*

[Link]
NamePolicy=kernel database onboard slot path
MACAddressPolicy=none
```

### Problem: Messages in /var/log/syslog about Network-Manager
### Fix:

    1. sudo snap remove network-manager
    - this reduced the amount of garbage in syslog when the containers fired up
    - I no longer saw the interfaces bouncing

### Problem: systemd udevd Could not generate persistent MAC address for veth*
### Possible Fix: (Reverted. It didn't work)
```
cat << EOF > /etc/udev/rules.d/01-net-setup-link.rules
SUBSYSTEM=="net", ACTION=="add|change", ENV{INTERFACE}==br-*", PROGRAM="/bin/sleep 0.5"
SUBSYSTEM=="net", ACTION=="add|change", ENV{INTERFACE}==docker*", PROGRAM="/bin/sleep 0.5"
EOF
```
    - this didn't do the trick
### Fix:
```
cat << EOF > /etc/systemd/network/99-default.link
[Link]
NamePolicy=kernel database onboard slot path
MACAddressPolicy=none
EOF
```


### Problem: Error about IPV6 connections failing
### Fix: Disable IPV6
    - In /etc/sysctl.conf add:
    ```
    net.ipv6.conf.all.disable_ipv6=1
    net.ipv6.conf.default.disable_ipv6=1
    net.ipv6.conf.lo.disable_ipv6=1
    net.ipv6.conf.wlan0.disable_ipv6=1
    ```
    - anddd sysctl -p


### Configure Docker MTU to be 1400 - We had issues with standard 1500 MTU
### Fix:
    1. Configure Docker Default Bridge
    ```
    cat << EOF > /etc/systemd/system/docker.service/override.conf
    [Service]
    ExecStart=/usr/bin/dockerd --mtu=1400 -H fd:// --containerd=/run/containerd/containerd.sock
    EOF
    ```
    
    2. Configure MTU for robot docker-compose. Append this to docker-compose.yaml
    ```
    networks:
        ros:
            driver: bridge
            driver_opts:
                com.docker.network.driver.mtu: 1400
    ```

### Problem: Video Stream over LTE craps out after above fixes
### Fix: Reduce bandwitdh requirments

    1. Enable compression on the ssh tunnel
    - set compression to yes in /etc/ssh/sshd_config

    2. Lower Camera Resolution
    - dropped it to 320 by 240








