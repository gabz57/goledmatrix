all: ledmatrix/build
.PHONY: ledmatrix/build

ledmatrix/build:
	@docker buildx build . \
	--platform linux/arm/v7 \
	--tag gabz57/goledmatrix:latest \
	--output bin/

#ledmatrix/build:
#	@docker buildx build . \
#	--platform linux/arm64 \
#	--tag gabz57/goledmatrix:latest \
#	--output bin/

#	--target bin \
#
#ledmatrix/build:
#	@docker buildx build . \
#	--platform linux/amd64,linux/arm64 \
#	--tag gabz57/goledmatrix:latest \
#	--target bin \
#	--output bin/

ledmatrix/push:
	@docker buildx build . \
	--platform linux/amd64,linux/arm64 \
	--tag gabz57/goledmatrix:latest \
	--push
