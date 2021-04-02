###############
# Build stage for linux/arm/v7 platform
FROM --platform=${BUILDPLATFORM} dockcross/linux-armv7 AS builder

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
WORKDIR /go/src/github.com/gabz57/goledmatrix/demo/_local
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 go build -o /out/example .


###############
# Running stage
FROM arm32v7/python:3.9.2-slim-buster AS bin
## TODO? COPY --from=builder # compiled C library
COPY --from=builder /out/example /usr/bin/goledmatrix
COPY ./resetmatrix.py .
COPY ./entrypoint.sh .
ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]
#CMD [ "/usr/bin/goledmatrix" ]