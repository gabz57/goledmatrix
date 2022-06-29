#!/bin/sh
set -e

docker run --rm \
        --pull=always \
        --privileged \
        --net=host \
        -e SERVER=1 \
        --cap-add=SYS_ADMIN \
        --cap-add=NET_ADMIN \
        -p 8080:8080 \
        --volume /var/run/dbus/:/var/run/dbus/:z \
        -it gabz57/goledmatrix:rpi64-server