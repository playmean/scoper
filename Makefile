GOOS := $(go env GOOS)
GOARCH := $(go env GOARCH)

bin::
	[ -d ./dist ] || mkdir -p ./dist
	CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -ldflags '-w -extldflags "-static"' -o main
	mv ./main ./dist/scoper

docker:
	docker build -t scoper -f ./Dockerfile ./dist

.PHONY: docker
