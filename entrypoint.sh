#!/bin/bash

#systemctl start hciuart
##service bluetooth stop
##service dbus stop
#service dbus start
#service dbus status
#hciconfig hci0
#service bluetooth start
#bluetoothd &
#
## wait for startup of services
#msg="Waiting for services to start..."
#time=0
#echo -n $msg
#while [[ "$(pidof start-stop-daemon)" != "" ]]; do
#    sleep 1
#    time=$((time + 1))
#    echo -en "\r$msg $time s"
#done
#echo -e "\r$msg done! (in $time s)"
#
#echo -n "Restarting hci0 ..."
#hciconfig hci0 down
#hciconfig hci0 up
#hciconfig hci0
#echo -n "Restarting hci0 DONE"
##hciconfig hci0 piscan
##hciconfig hci0 sspmode 1
#
#echo -n "Pairing (1/2) ..."
#bluetoothctl -- list
#bluetoothctl -- power on
#bluetoothctl -- default-agent
#bluetoothctl -- agent on
#bluetoothctl -- scan on
#echo -n "Scanning 30s ..."
#bluetoothctl -- trust F4:93:9F:C0:64:D4
#sleep 30
#echo -n "Pairing (2/2) ..."
#bluetoothctl -- pair F4:93:9F:C0:64:D4
#bluetoothctl -- connect F4:93:9F:C0:64:D4
#echo -n "Stop scanning ..."
#bluetoothctl -- scan off
#
#echo -n "Waiting 5s ..."
#sleep 5
#
echo -n "Running matrix..."
python3 resetmatrix.py
cd /usr/bin/
./goledmatrix
