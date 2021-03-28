all: ledmatrix/build
.PHONY: ledmatrix/build

ledmatrix/build:
	@docker buildx build . \
	--platform linux/amd64,linux/arm/v7 \
	--tag gabz57/goledmatrix:core
#	--target bin/ledmatrix \
#	--output bin/ledmatrix/

ledmatrix/push:
	@docker buildx build . \
	--platform linux/amd64,linux/arm/v7 \
	--tag gabz57/goledmatrix:core \
	--push
#	--target bin \
#	--output bin/ \
#
#rpc/build:
#	@docker buildx build rpc/. \
#	--platform linux/amd64,linux/arm/v7 \
#	--tag gabz57/goledmatrix:rpc
##	--target bin/ledmatrix \
##	--output bin/ledmatrix/
#
#rpc/push:
#	@docker buildx build rpc/. \
#	--platform linux/amd64,linux/arm/v7 \
#	--tag gabz57/goledmatrix:rpc \
#	--push
##	--target bin \
##	--output bin/ \