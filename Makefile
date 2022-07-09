all: ledmatrix64/push
.PHONY: ledmatrix64/push

ledmatrix64/build:
	@docker buildx build . \
	--platform linux/arm64 \
	--tag gabz57/goledmatrix:rpi64 \
	--output bin/64/
# docker buildx build . --platform linux/arm64 --tag gabz57/goledmatrix:rpi64 --output bin/64/

ledmatrix64/push:
	@docker buildx build . \
	--platform linux/arm64 \
	--tag gabz57/goledmatrix:rpi64 \
	--push
# docker buildx build . --platform linux/arm64 --tag gabz57/goledmatrix:rpi64 --push
