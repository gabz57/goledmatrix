# Go Led Matrix

**Many thanks to Henner Zeller for its original library : [rpi-rgb-led-matrix](https://github.com/hzeller/rpi-rgb-led-matrix) & 
and to Máximo Cuadros for its [Go binding](https://github.com/mcuadros/go-rpi-rgb-led-matrix) from wich this code is inspired**

I decided to start a new project instead of forking Máximo Cuadros project as it required some breaking changes,
and the need to handle the Emulator UI on the main thread for MacOsX laptop (at least, I also wanted to learn new things starting from scratch)

I also wanted to include some tooling for working/building locally and simply update 
and run the code from a fresh RPi without having to restart its setup each time, which takes some time and is error prone.

Starting the hardware matrix also requires to init the panels with a python script

## For developers

This project uses Docker BuildX to build and prepare the different targets :
- **linux/amd64** (for MacBook)
- **linux/aarch64** (for RPi 3B+)
// note I'm note sure about this yet, for RPi 3B+ and GPIO arm/v7 might be needed, arm64 should be written instead of aarch64

### To do only once : Setup Docker BuildX (⚠️ One must also ensure that experimental mode is enabled in Docker)
```sh
$ docker run --privileged --rm tonistiigi/binfmt --install all
$ docker buildx create --use --name rpibuilder
$ docker buildx inspect --bootstrap
```

### Build using Docker BuildX
```sh
$ docker buildx build --platform linux/amd64,linux/arm/v7 .
```

### Local manual GO build (emulator only)
```sh
$ cd demo/_local
$ go mod vendor
$ go build -o ./out/example .
```

### Build & publish to Docker HUB
Instead of typing long docker commands, a Makefile is available to run these commands

The Makefile default behaviour only builds the project, if you wish to publish it to Dockerhub, you must first be logged into Docker, and update the command to publish it to your Docker Hub repository

```sh
$ make
OR 
# this will create a local executable file for RPi
$ docker buildx build --platform linux/arm/v7 . --output bin/ledmatrix/
# this will create a local executable file for MacBook
$ docker buildx build --platform linux/amd64 . --output bin/ledmatrix/
```

### Run on MacBook
I can not run it yet on my own macbook via docker run command (while compiling C code, Go code in standalone with emulator works fine), as I'm having trouble at running a X11 server with the current security policy, and cannot bypass it...
Some help at trying to solve this can be found [here](https://gist.github.com/cschiewek/246a244ba23da8b9f0e7b11a68bf3285)
```sh
$ docker run -ti --rm -e DISPLAY=$DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix -v $HOME/.Xauthority:$HOME/.Xauthority gabz57/goledmatrix:demo
$ docker run -ti --rm -e DISPLAY=host.docker.internal:0 gabz57/goledmatrix:demo
```

## On the Raspberry Pi

TBD

### Setup Docker

TBD

## Try your own code

See `/demo directory for a few examples

To start a new matrix, you can use this template and run it with MATRIX_EMULATOR=1 environment variable

```go
package main

import (
    . "github.com/gabz57/goledmatrix"
)

func main() {
    RunMatrices(app)
}

func app() {
    Run(func(config *MatrixConfig) (Matrix, error) {
        return BuildMatrix(config)
    }, Gameloop)
}

func Gameloop(c *Canvas, done chan struct{}) {
    // manipulate canvas here, call c.Render() to publish the canvas to the matrix (Hardware or emulator)
}
```

# A few notes

What I try to achieve is writing a small library over the Go binding to create and manipulate some UI components, and being able to prepare and render them on my laptop before pushing the code to the hardware.

What I call hardware here is my small setup using a RPi 3B, and 4 64*64 leds panels.

Before the go version, I tried the `C++` version, which I found quite hard to read and maintain, I also miss experience on writing test in this langage, even if I really appreciate the low level control.

I also tried and had a fully working `TypeScript` version, but had some difficulties to write an emulator using either a web page, contact me if you have hints on this ;)

I finally tried the `Go` binding, the emulator worked with a few changes, but the actual binding doesn't seem actively maintained.
There are other difficulties as I'm working on a machine on which I don't want/need to install gcc or other C related compilers,
using Docker to handle the build should allow me to prepare a runnable image for my RPi, but plugin the C related code in Go prevents me to run this code and test the emulator. Any suggestion would be appreciated here ;)

We should expect 3 possibles setups :
- Emulator only on any machine (run via IDE, via building & running manually the Go code, or by building & running the docker image if we can pass $DISPLAY)
- RPi only -> use docker run ... directly on the RPi
- on LAN: RPi server + other machine as client (can be your IDE, with or without an emulator) -> the RPi open a RPC server to allow remote control of the canvas, which is then applied to the hardware



# WIP ZONE

## ℹ️ IMPORTANT

The emulator works on MacBook, but I canno't yet compile the whole code with dependency to the C library, one need to comment the file and switch case for NewRGBLedMatrix in BuildMatrix method.

IDEAL : We should be able to build an image using GOARCH=arm/v7 or GOARCH=arm64 and CGO_ENABLED=1, using Docker which can also embed QEMU to allow running code compiled for another target directly by running the docker image

```sh

# What I want to use for building the Go app for ARM32 (later via Docker), something like (but for 32-bit OS)
$ CGO_ENABLED=1 CC=???? GOOS=linux GOARCH=arm go build -o ./out/example/ demo.go

# Next line show how I would like to build for MacBook (AMD64) (only  for starting the emulator or server part, 
# thus we could exclude the matrix_rpi.go file from compilation, which references the libmatrix C library)
$ CGO_ENABLED=0 CC=gcc GOOS=linux GOARCH=amd64 go build -o ./out/example/ demo.go
```

Currently, AFTER commenting Go matrix_rpi.go file, one can build using docker toward ARM64 with CGO_ENABLED=0

But we won't be able to run it for hardware tests, thus losing any interest...

Using :
```sh
$ docker buildx build --platform linux/arm64 . --output bin/ledmatrix/
# which will run in Dockerfile
# RUN CGO_ENABLED=0 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o /out/example/ .
```

When trying to turn on C compilation with Go application (CGO_ENABLED=1)
```sh
$ docker buildx build --platform linux/arm64 . --output bin/ledmatrix/
# which will run in Dockerfile
# RUN CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o /out/example/ .

# OUTPUT
 => ERROR [builder 8/8] RUN CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o /out/example/ .                                                                                            20.1s 
------                                                                                                                                                                                                                  
 > [builder 8/8] RUN CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o /out/example/ .:                                                                                                        
#12 19.85 # github.com/gabz57/goledmatrix
#12 19.85 /usr/lib/gcc-cross/aarch64-linux-gnu/6/../../../../aarch64-linux-gnu/bin/ld: skipping incompatible ../../vendor/rpi-rgb-led-matrix/lib/librgbmatrix.a when searching for -lrgbmatrix
#12 19.85 /usr/lib/gcc-cross/aarch64-linux-gnu/6/../../../../aarch64-linux-gnu/bin/ld: cannot find -lrgbmatrix
#12 19.85 collect2: error: ld returned 1 exit status

```
[Resource qui distingue arm/v7 & arm64, et tuto pour installer docker sur RPi](https://withblue.ink/2020/06/24/docker-and-docker-compose-on-raspberry-pi-os.html)

In the Docker ecosystem, 
64-bit ARM images are called arm64 or arm64/v8.
- Raspberry 3B+ si OS en 64-bit (récent, pas certain que la compatibilité GPIO soit totale)
- Raspberry 4

32-bit ARM images for Raspberry Pi OS are labeled as 
armhf, armv7, or arm/v7
- Raspberry 2, 3 si OS en 32-bit

[cross compilation Go](https://connect.ed-diamond.com/GNU-Linux-Magazine/GLMFHS-106/Utiliser-simplement-un-reseau-de-neurones-sur-Raspberry-Pi-grace-a-ONNX-et-Go)

[buildX multiarch](https://medium.com/nttlabs/buildx-multiarch-2c6c2df00ca2)