bin::
	[ -d ./dist ] || mkdir -p ./dist
	CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-s -w -extldflags "-static"' -o ./dist/scoper

docker:
	docker build -t playmean/scoper -f ./Dockerfile .

.PHONY: bin docker
