######################################
# Build C lib for linux/arm64 platform
FROM --platform=${BUILDPLATFORM} dockcross/linux-arm64 AS cbuilder

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
FROM --platform=${BUILDPLATFORM} dockcross/linux-arm64 AS gobuilder
RUN wget https://dl.google.com/go/go1.16.3.linux-arm64.tar.gz
RUN tar -xvf go1.16.3.linux-arm64.tar.gz
RUN mv go /usr/local
ENV GOROOT /usr/local/go
ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

COPY --from=cbuilder /c/rpi-rgb-led-matrix /go/src/github.com/gabz57/goledmatrix/vendor/rpi-rgb-led-matrix
COPY ./. /go/src/github.com/gabz57/goledmatrix/
# overwrite BuildMatrix method with Hardware binding
COPY ./matrix_rpi /go/src/github.com/gabz57/goledmatrix/matrix_rpi.go
COPY ./matrix_builder_rpi /go/src/github.com/gabz57/goledmatrix/matrix_builder.go

## build Go DEMO application
WORKDIR /go/src/github.com/gabz57/goledmatrix/demo
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o /out/goledmatrix-bin .

###############
# Running stage
FROM arm64v8/python:3.9.2-slim-buster AS bin
RUN pip3 install gpiozero
COPY ./fonts /usr/bin/fonts
COPY ./resetmatrix.py .
COPY ./entrypoint.sh .

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]

EXPOSE 8080

COPY --from=gobuilder /out/goledmatrix-bin /usr/bin/goledmatrix
