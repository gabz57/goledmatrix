######################################
# Build C lib for linux/arm64 platform
FROM --platform=${BUILDPLATFORM} dockcross/linux-arm64-lts AS cbuilder

## To drive hardware matrix via GPIO on RPi
## fetch origial C library via Git submodule & build it
RUN apt-get update && apt-get install -y git

WORKDIR /c/
RUN git clone https://github.com/hzeller/rpi-rgb-led-matrix.git
WORKDIR /c/rpi-rgb-led-matrix
## Note: only building the library for librgbmatrix.a file (skipping samples which makes compilation fail)
RUN make -C ./lib

#################
# Building GO app
FROM --platform=${BUILDPLATFORM} dockcross/linux-arm64-lts AS gobuilder
RUN apt-get update && apt-get install -y golang

COPY --from=cbuilder /c/rpi-rgb-led-matrix /go/src/github.com/gabz57/goledmatrix/vendor/rpi-rgb-led-matrix
COPY ./. /go/src/github.com/gabz57/goledmatrix/

## build Go application
WORKDIR /go/src/github.com/gabz57/goledmatrix
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o /out/goledmatrix-bin .

###############
# Running stage
FROM arm64v8/python:3.9.12-slim-buster AS bin

RUN apt-get update && apt-get install -y bluez dbus bluetooth

RUN pip3 install gpiozero

COPY canvas/fonts /usr/bin/canvas/fonts
COPY ./img /usr/bin/img
COPY ./resetmatrix.py .
COPY ./entrypoint.sh .

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]

EXPOSE 8080

COPY --from=gobuilder /out/goledmatrix-bin /usr/bin/goledmatrix
