FROM --platform=${BUILDPLATFORM} golang:1.16.2-stretch AS builder
ARG TARGETOS
ARG TARGETARCH

RUN apt-get update && apt-get install -y gcc-aarch64-linux-gnu g++-aarch64-linux-gnu

ADD . /go/src/github.com/gabz57/goledmatrix


## fetch & build C library to drive hardware matrix via GPIO on RPi
WORKDIR /go/src/github.com/gabz57/goledmatrix/vendor/rpi-rgb-led-matrix
RUN git submodule update --init
RUN make

#WORKDIR app/vendor/rpi-rgb-led-matrix
## build Go DEMO application
#WORKDIR /go/src/github.com/gabz57/goledmatrix
WORKDIR /go/src/github.com/gabz57/goledmatrix/demo/_local

RUN CGO_ENABLED=0 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o /out/example/ .
#RUN CGO_ENABLED=0 CC=aarch64-linux-gnu-gcc GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/example .

# Final stage - pick any old arm64 image you want
#FROM multiarch/ubuntu-core:arm64-bionic
FROM scratch

COPY --from=builder /out/example /usr/bin/goledmatrix
CMD [ "/usr/bin/goledmatrix" ]