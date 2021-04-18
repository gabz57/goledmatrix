############### @DEPRECATED: I switched to arm64 on my Rpi3B+
# Build stage for linux/arm/v7 platform
FROM --platform=${BUILDPLATFORM} dockcross/linux-armv7 AS builder

RUN apt-get update && apt-get install -y git
RUN wget https://dl.google.com/go/go1.16.3.linux-armv6l.tar.gz
RUN tar -xvf go1.16.3.linux-armv6l.tar.gz
RUN mv go /usr/local
ENV GOROOT /usr/local/go
ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

COPY ./. /go/src/github.com/gabz57/goledmatrix
# overwrite BuildMatrix method with Hardware binding
COPY ./matrix/matrix_rpi /go/src/github.com/gabz57/goledmatrix/matrix/matrix_rpi.go
COPY ./matrix/matrix_builder_rpi /go/src/github.com/gabz57/goledmatrix/matrix/matrix_builder.go

## To drive hardware matrix via GPIO on RPi
## fetch origial C library via Git submodule & build it
WORKDIR /go/src/github.com/gabz57/goledmatrix/vendor/rpi-rgb-led-matrix
RUN git submodule update --init
## Note: only building the library for librgbmatrix.a file (skipping samples which makes compilation fail)
RUN make -C ./lib

## build Go application
WORKDIR /go/src/github.com/gabz57/goledmatrix
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 go build -o /out/example .


###############
# Running stage
FROM arm32v7/python:3.9.2-slim-buster AS bin
RUN pip3 install gpiozero
COPY ./fonts /usr/bin/fonts
COPY ./img /usr/bin/img
COPY ./resetmatrix.py .
COPY ./entrypoint.sh .

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]

EXPOSE 8080

COPY --from=builder /out/example /usr/bin/goledmatrix
