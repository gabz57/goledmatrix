all: ledmatrix64/push
.PHONY: ledmatrix64/push

#ledmatrix32/build:
#	@docker buildx build . -f armv7.Dockerfile \
#	--platform linux/arm/v7 \
#	--tag gabz57/goledmatrix:rpi32 \
#	--output bin/32/
#
#ledmatrix32/push:
#	@docker buildx build . -f armv7.Dockerfile \
#	--platform linux/arm/v7 \
#	--tag gabz57/goledmatrix:rpi32 \
#	--push

ledmatrix64/build:
	@docker buildx build . -f arm64.Dockerfile \
	--platform linux/arm64 \
	--tag gabz57/goledmatrix:rpi64 \
	--output bin/64/

ledmatrix64/push:
	@docker buildx build . -f arm64.Dockerfile \
	--platform linux/arm64 \
	--tag gabz57/goledmatrix:rpi64 \
	--push
