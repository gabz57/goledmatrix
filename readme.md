# Go Led Matrix

**Many thanks to Henner Zeller for its original library : [rpi-rgb-led-matrix](https://github.com/hzeller/rpi-rgb-led-matrix) & 
and to Máximo Cuadros for its [Go binding](https://github.com/mcuadros/go-rpi-rgb-led-matrix) from wich this code is inspired.**

I decided to start a new project instead of forking Máximo Cuadros project as it required some breaking changes,
and the need to handle the Emulator UI on the main thread for MacOsX laptop (at least, I also wanted to learn new things starting from scratch).

I also wanted to include some tooling for working/building locally and simply update 
and run the code from a fresh RPi without having to restart its setup each time, which takes some time and is error prone.

Starting the hardware matrix also requires to be intialized with a script written in python.

## Build

If you only need to use the emulator, you can skip the build part using Docker and go to the **RUN** section.


This project uses Docker BuildX to build and prepare the different targets :
- **linux/arm/v7** (for RPi 3B+ with previous OS 32-bit)
- **linux/arm64** (for RPi 3B+ with more recent OS in 64-bit)
// Note arm64 might not work and need further tests


MacBook:
```sh
# https://gobot.io/documentation/platforms/joystick/
brew install sdl2
# https://github.com/gopherdata/gophernotes/issues/82
brew install pkg-config
```

### To do only once : 
Setup Docker BuildX (⚠️ One must also ensure that experimental mode is enabled in Docker)
```sh
$ docker run --privileged --rm tonistiigi/binfmt --install all
$ docker buildx create --use --name rpibuilder
$ docker buildx inspect --bootstrap
```

### Build using Docker BuildX
To ensure the build is working, run it via a make command shortcut :
```sh
$ make
# OR (equivalent to)
$ make ledmatrix64/build
```
To push the code to Dockerhub, one can do the same
(you must be authenticated or push to your own repository)
```sh
$ make ledmatrix64/push
```
There is also an unmaintained 32 bits version (for <= RPi3 boards, or RPi3 with previous 32bits OS)

## Run on any computer (locally using emulator)

This applications provides files that we be used during Docker build to plug in c library, 
these files are not used by default for development for being able to run the emulator without additional operation.

### Local GO build (for emulator only, without using C rpi-rgb-led-matrix library)
One can run manually the code with the following commands
```sh
$ go mod vendor
$ cd demo
$ go build -o ./out/example .
```
### Using IDE 😎
ℹ️: Any IDE can also run this code, with IntelliJ you simply have to click on play inside the main.go file.
Do not forget to configure ENV variables to enable the emulator (`MATRIX_EMULATOR=1`)

[comment]: <> (### Build & publish to Docker HUB)

[comment]: <> (Instead of typing long docker commands, a Makefile is available to run these commands)

[comment]: <> (The Makefile default behaviour only builds the project, if you wish to publish it to Dockerhub, you must first be logged into Docker, and update the command to publish it to your Docker Hub repository)

[comment]: <> (```sh)

[comment]: <> ($ make)

[comment]: <> (OR )

[comment]: <> (# this will create a local executable file for RPi 32-bit)

[comment]: <> ($ docker buildx build --platform linux/arm/v7 . -f armv7.Dockerfile --output bin/ledmatrix/)

[comment]: <> (# this will create a local executable file for RPi 64-bit)

[comment]: <> ($ docker buildx build --platform linux/arm64 . -f arm64.Dockerfile --output bin/ledmatrix/)

[comment]: <> (```)

### [KO] Run the Raspbian 64-bits OS version on MacBook
I could not run it yet on my own macbook via docker run command (while compiling C code, Go code in standalone with emulator works fine), as I'm having trouble at running a X11 server with the current security policy, and cannot bypass it...

Some help at trying to solve this can be found [here](https://gist.github.com/cschiewek/246a244ba23da8b9f0e7b11a68bf3285)
```sh
$ docker run -ti --rm -e DISPLAY=$DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix -v $HOME/.Xauthority:$HOME/.Xauthority gabz57/goledmatrix:demo
$ docker run -ti --rm -e DISPLAY=host.docker.internal:0 gabz57/goledmatrix:demo
```

## Run on the Raspberry Pi

One only need Docker to start the Go Led Matrix application.  
Note that the GPIO lib requires docker to be run with `--privileged` for running as root

### Setup Docker

TBD

### Run the app !!

To run the latest version of the application on your RPi, simply run :
```sh
$ docker run --rm --privileged gabz57/goledmatrix:rpi64
```
To run the application as a server controlled by a remote instance, just enable the server mode :
```sh
$ docker run --rm --privileged -e MATRIX_SERVER=1 -p 8080 gabz57/goledmatrix:rpi64
```
Then control it from another instance started as client :
```sh
$ docker run --rm -e MATRIX_CLIENT=1 -e MATRIX_ADDRESS=192.168.1.14 -p 8080 gabz57/goledmatrix:rpi64
```

Cleaning the RPi:
```sh
$ docker ps
$ docker stop CONTAINER
$ docker rmi gabz57/goledmatrix:rpi64
```

## Try your own code

See `/demo directory for a few examples

To start a new matrix, you can use this template and run it with MATRIX_EMULATOR=1 environment variable

```go
package main

import (
	"time"

	. "github.com/gabz57/goledmatrix"
	. "github.com/gabz57/goledmatrix/components"
)

func main() {
    RunMatrices(goLedApplication)
}

func goLedApplication() {
    Run(Gameloop)
}

func Gameloop(c Canvas, done chan struct{}) {
    // Your code starts here
    // Example (note: declaring scene with duration might change in a close future, 
    // allowing to run a single scene without duration, 
    // and changing the way they end
    sceneDuration := 12 * time.Second
    engine := NewEngine(c, []*Scene{
        NewScene([]Component{myComponent(*c)}, sceneDuration),
    })
    engine.Run(done)
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

We can find 3 possibles setups :
- Emulator only on any machine (run via IDE, via building & running manually the Go code, or by building & running the docker image if we can pass $DISPLAY)
- RPi only -> use docker run ... directly on the RPi
- on LAN: one RPi acts as server + another machine as client (can be your IDE, or another RPi) -> the server listen for RPC calls to allow remote control of the canvas, which is then applied to the hardware


# WIP ZONE

## ℹ️ IMPORTANT

The emulator works on MacBook, but I cannot compile the whole code with its dependency to the C library.
One need to comment the file and switch off the case for NewRGBLedMatrix in BuildMatrix method.

IDEAL : We should be able to build an image using GOARCH=arm/v7 or GOARCH=arm64 and CGO_ENABLED=1.
This image should run on macbook (using emulator or server mode) using Docker which can embed QEMU 
QEMU allow compiled code (for another target) to run directly on MacOS.

`HELP NEEDED` using the emulator through the docker image requires to provide some kind of display to docker run command.
This should looks like :
```sh
# these are just examples, they don't work
$ docker run -ti --rm -e DISPLAY=$DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix -v $HOME/.Xauthority:$HOME/.Xauthority gabz57/goledmatrix:demo
$ docker run -ti --rm -e DISPLAY=host.docker.internal:0 gabz57/goledmatrix:demo
```

On MacBook, when trying to turn on C compilation with Go application (CGO_ENABLED=1) using previous Dockerfile version :
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

## Links
####cross compilation in Go
https://connect.ed-diamond.com/GNU-Linux-Magazine/GLMFHS-106/Utiliser-simplement-un-reseau-de-neurones-sur-Raspberry-Pi-grace-a-ONNX-et-Go

#### buildX multiarch
https://medium.com/nttlabs/buildx-multiarch-2c6c2df00ca2

#### Building its own toolchain
https://rolandsdev.blog/cross-compile-for-raspberry-pi-with-docker/
#### GoReleaser + golang-cross
https://goreleaser.com/limitations/cgo/
https://goreleaser.com/cookbooks/cgo-and-crosscompiling

https://github.com/troian/golang-cross


## Trying again, using 2 Dockefile for arm32 & arm64
```sh
$ docker buildx build --platform linux/arm/v7 . -f armv7.Dockerfile --output bin/ledmatrix/
$ docker buildx build --platform linux/arm64 . -f arm64.Dockerfile --output bin/ledmatrix/
```

## Build all locally ℹ️ using 2 Dockerfile
```sh
$ make ledmatrix32/build
$ make ledmatrix64/build
$ file bin/32/usr/bin/goledmatrix
# bin/32/usr/bin/goledmatrix: ELF 32-bit LSB executable, ARM, EABI5 version 1 (SYSV), dynamically linked, interpreter /lib/ld-linux-armhf.so.3, for GNU/Linux 4.10.8, Go BuildID=e64a86f74eb3292dac7d89cb25d93e9f58fef28b, with debug_info, not stripped
$ file bin/64/usr/bin/goledmatrix
# bin/64/usr/bin/goledmatrix: ELF 64-bit LSB executable, ARM aarch64, version 1 (SYSV), dynamically linked, interpreter /lib/ld-linux-aarch64.so.1, for GNU/Linux 4.10.8, Go BuildID=0e5708a60115cc43d38506ffdc1fa1ad81e7719a, with debug_info, not stripped
```
