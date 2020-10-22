bin::
	[ -d ./dist ] || mkdir -p ./dist
	CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o ./dist/scoper

docker:
	docker build -t scoper -f ./Dockerfile ./dist

start:
	docker run --rm -d -p 3000:3000/tcp -v "`pwd`/dist/data:/data" scoper:latest

.PHONY: docker start
