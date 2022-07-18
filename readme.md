# Go Led Matrix

**Many thanks to Henner Zeller for its original library : [rpi-rgb-led-matrix](https://github.com/hzeller/rpi-rgb-led-matrix) & 
and to Máximo Cuadros for its [Go binding](https://github.com/mcuadros/go-rpi-rgb-led-matrix) from wich this code is based on.**

I decided to start a new project instead of forking Máximo Cuadros project with the need to handle the Emulator UI on the main thread for MacOsX laptop (at least, I also wanted to learn new things starting from scratch).

I also wanted to include some tooling for working/building locally and simply update and run the code from a fresh RPi without having to restart its setup each time, which takes some time and is error-prone.

Starting the hardware matrix also requires to be initialized with a script given in Python (handled in dockerfile).

## Build

If you only need to use the emulator, you can skip the build part using Docker and go to the **RUN** section.

This project uses Docker `buildX` to build and prepare the `Raspberry` targets :
- **linux/arm64** (for RPi 3B+ with more recent OS in 64-bit)

### To be done only once : 
Setup Docker BuildX (⚠️ One must also enable experimental mode in Docker)
```sh
$ docker run --privileged --rm tonistiigi/binfmt --install all
$ docker buildx create --use --name rpibuilder
$ docker buildx inspect --bootstrap
```

On MacBook, one may require to install libraries relative to gamepad usage via bluetooth:
```sh
# https://gobot.io/documentation/platforms/joystick/
brew install sdl2
# https://github.com/gopherdata/gophernotes/issues/82
brew install pkg-config
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

## Run on Macbook

This applications provides files that we be used during Docker build to plug in c library, 
these files are not used by default for development for being able to run the emulator without additional operation.

### Local GO build (for emulator only, without using C rpi-rgb-led-matrix library)
One can run manually the code with the following commands
```sh
$ go mod vendor
$ cd demo
$ go build -o ./out/example .
```
### Using IDE (IntelliJ 😎)
ℹ️: Any IDE can also run this code, with IntelliJ you simply have to click on play inside the main.go file.
Do not forget to configure ENV variables to enable the emulator (`MATRIX_EMULATOR=1`)

## Run on the Raspberry Pi

Only Docker is needed to start the GoLedMatrix application.  
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

* Note that `goled.sh` and `goserverled.sh` scripts can be copied to the raspberry and avoid you to typing these commands, they also contain some bluetooth settings and tests 

Cleaning the RPi:
```sh
$ docker ps
$ docker stop CONTAINER
$ docker rmi gabz57/goledmatrix:rpi64
```

## Demo !

Displaying audio frequencies (from laptop microphone)

https://user-images.githubusercontent.com/3730276/179513887-a3f979a3-a2d1-4764-8f10-85715c25532f.mov

Fadings randomly colored dots

https://user-images.githubusercontent.com/3730276/179515562-23323623-9f0b-4cd0-aa39-615a35fdc24a.mov

Clock

https://user-images.githubusercontent.com/3730276/179516728-51b4286a-f51f-459e-9189-743ca0a58675.mov

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

Before the go version, I tried the `C++` version, which I found quite hard to read and maintain, I also miss experience on writing tests in this language, even if I really appreciate the low level control.

I also tried and had a fully working `TypeScript` version, but had some difficulties to write an emulator using either a web page, contact me if you have hints on this ;)

I finally tried the `Go` binding, the emulator worked fine with a few changes.

We can find 3 possibles setups :
- `Emulator only` on any machine (at least MacBook via IDE, or via building & running manually the Go code)
- `RPi only` -> using `docker run ...` or scripts directly on the RPi
- `Remote mode`: on LAN, a Raspberry acts as server and is plugged to the panel + another machine as client (can be your IDE, or another RPi) -> the server listen for RPC calls to allow remote control of the canvas, which is then simply applied to the hardware

## Useful links
#### Go cross compilation
https://github.com/troian/golang-cross

https://connect.ed-diamond.com/GNU-Linux-Magazine/GLMFHS-106/Utiliser-simplement-un-reseau-de-neurones-sur-Raspberry-Pi-grace-a-ONNX-et-Go

#### buildX multiarch
https://medium.com/nttlabs/buildx-multiarch-2c6c2df00ca2

#### Building its own toolchain
https://rolandsdev.blog/cross-compile-for-raspberry-pi-with-docker/
#### GoReleaser + golang-cross
https://goreleaser.com/limitations/cgo/

https://goreleaser.com/cookbooks/cgo-and-crosscompiling
