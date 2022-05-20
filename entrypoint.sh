#!/bin/bash

#service bluetooth stop
#service dbus stop
service dbus start
service bluetooth start
#bluetoothd &

# wait for startup of services
msg="Waiting for services to start..."
time=0
echo -n $msg
while [[ "$(pidof start-stop-daemon)" != "" ]]; do
    sleep 1
    time=$((time + 1))
    echo -en "\r$msg $time s"
done
echo -e "\r$msg done! (in $time s)"

hciconfig hci0 down
hciconfig hci0 up
#hciconfig hci0 piscan
#hciconfig hci0 sspmode 1

msg="Pairing (1/2) ..."
echo -n $msg
bluetoothctl -- power on
bluetoothctl -- agent on
bluetoothctl -- default-agent
bluetoothctl -- scan on &
msg="Waiting 20s ..."
echo -n $msg
sleep 20
msg="DONE Waiting 20s ..."
echo -n $msg
bluetoothctl -- scan off
msg="Pairing (2/2) ..."
echo -n $msg
bluetoothctl -- pair F4:93:9F:C0:64:D4
bluetoothctl -- trust F4:93:9F:C0:64:D4
bluetoothctl -- connect F4:93:9F:C0:64:D4
sleep 2
msg="Connected..."
echo -n $msg

python3 resetmatrix.py
cd /usr/bin/
./goledmatrix
