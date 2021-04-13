all: ledmatrix/build
.PHONY: ledmatrix/build

ledmatrix/build:
	@docker build . -f Dockerfile \
	--tag gabz57/goledmatrix:rpi32 \
	--output bin/32/

ledmatrix32/build:
	@docker buildx build . -f armv7.Dockerfile \
	--platform linux/arm/v7 \
	--tag gabz57/goledmatrix:rpi32 \
	--output bin/32/

ledmatrix32/push:
	@docker buildx build . -f armv7.Dockerfile \
	--platform linux/arm/v7 \
	--tag gabz57/goledmatrix:rpi32 \
	--push

ledmatrix64/build:
	@docker buildx build . -f arm64.demo.Dockerfile \
	--platform linux/arm64 \
	--tag gabz57/goledmatrix:rpi64 \
	--output bin/64/

ledmatrix64-demo/push:
	@docker buildx build . -f arm64.demo.Dockerfile \
	--platform linux/arm64 \
	--tag gabz57/goledmatrix:rpi64 \
	--push

ledmatrix64-server/push:
	@docker buildx build . -f arm64.server.Dockerfile \
	--platform linux/arm64 \
	--tag gabz57/goledmatrix:rpi64-server \
	--push
