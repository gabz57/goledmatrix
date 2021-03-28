FROM --platform=${BUILDPLATFORM} golang:1.16.2-alpine3.13 AS build
ARG TARGETOS
ARG TARGETARCH
#RUN apk update
#RUN apk upgrade
#RUN apk add --update go=1.8.3-r0 gcc=6.3.0-r4 g++=6.3.0-r4
WORKDIR /ledmatrix
ENV CGO_ENABLED=0
#ENV CGO_ENABLED=1 GOOS=linux // go install -a server
COPY . .
WORKDIR /demo/emulator
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/example .

FROM scratch AS bin-unix
COPY --from=build /out/example /usr/bin/goledmatrix

FROM bin-unix AS bin-linux
#FROM bin-unix AS bin-darwin
#FROM scratch AS bin-windows
#COPY --from=build /out/example /example.exe

FROM bin-${TARGETOS} AS bin
ENV MATRIX_EMULATOR=1
CMD [ "/usr/bin/goledmatrix" ]
