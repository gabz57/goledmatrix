###############
# Build stage for linux/arm64 platform
FROM --platform=${BUILDPLATFORM} dockcross/linux-arm64 AS builder

RUN apt-get update && apt-get install -y git
RUN wget https://dl.google.com/go/go1.16.3.linux-arm64.tar.gz
RUN tar -xvf go1.16.3.linux-arm64.tar.gz
RUN mv go /usr/local
ENV GOROOT /usr/local/go
ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

ADD . /go/src/github.com/gabz57/goledmatrix

## To drive hardware matrix via GPIO on RPi
## fetch origial C library via Git submodule & build it
WORKDIR /go/src/github.com/gabz57/goledmatrix/vendor/rpi-rgb-led-matrix
RUN git submodule update --init
## Note: only building the library for librgbmatrix.a file (skipping samples which makes compilation fail)
RUN make -C ./lib

## build Go DEMO application
WORKDIR /go/src/github.com/gabz57/goledmatrix/demo/_local
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o /out/example .

###############
# Running stage
FROM arm64v8/python:3.9.2-slim-buster AS bin
RUN pip3 install gpiozero
## TODO ? COPY --from=builder # compiled C library
COPY --from=builder /out/example /usr/bin/goledmatrix
COPY ./fonts /usr/bin/fonts
COPY ./resetmatrix.py .

COPY ./entrypoint.sh .

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]