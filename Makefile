all: ledmatrix32/build
.PHONY: ledmatrix32/build

ledmatrix32/build:
	@docker buildx build . -f armv7.Dockerfile \
	--platform linux/arm/v7 \
	--tag gabz57/goledmatrix:latest \
	--output bin/32/

ledmatrix32/push:
	@docker buildx build . -f armv7.Dockerfile \
	--platform linux/arm/v7 \
	--tag gabz57/goledmatrix:latest \
	--push

ledmatrix64/build:
	@docker buildx build . -f arm64.Dockerfile \
	--platform linux/arm64 \
	--tag gabz57/goledmatrix:latest \
	--output bin/64/

ledmatrix64/push:
	@docker buildx build . -f arm64.Dockerfile \
	--platform linux/arm64 \
	--tag gabz57/goledmatrix:latest \
	--push
