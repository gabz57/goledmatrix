###############
# Build stage for linux/arm/v7 platform
FROM dockcross/linux-armv7 AS builder

RUN apt-get update && apt-get install -y git golang
# TODO: describe why it works 😎 (inspired from this discussion: https://github.com/docker-library/golang/issues/129)
ENV GOPATH $HOME/go

ADD . /go/src/github.com/gabz57/goledmatrix

## To drive hardware matrix via GPIO on RPi
## fetch origial C library via Git submodule & build it
WORKDIR /go/src/github.com/gabz57/goledmatrix/vendor/rpi-rgb-led-matrix
RUN git submodule update --init
## Note: only building the library for librgbmatrix.a file (skipping samples which makes compilation fail)
RUN make -C ./lib

## build Go DEMO application
WORKDIR /go/src/github.com/gabz57/goledmatrix
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 go build -o /out/example .


###############
# Running stage
FROM arm32v7/python:3.9.2-slim-buster AS bin
RUN apt-get update \
 && apt-get install -y sudo

RUN adduser --disabled-password --gecos '' docker
RUN adduser docker sudo
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
RUN pip3 install gpiozero
## TODO ? COPY --from=builder # compiled C library
COPY --from=builder /out/example /usr/bin/goledmatrix
COPY ./fonts /usr/bin/fonts
COPY ./img /usr/bin/img
COPY ./resetmatrix.py .

COPY ./entrypoint.sh .

USER docker
ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]