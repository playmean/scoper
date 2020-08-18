bin::
	[ -d ./dist ] || mkdir -p ./dist
	CGO_ENABLED=0 go build -a -ldflags '-w -extldflags "-static"' -o main
	mv ./main ./dist/error-tracking

docker:
	docker build -t error-tracking -f ./Dockerfile ./dist

restart:
	docker-compose -f ./docker-compose.yml down && docker-compose -f ./docker-compose.yml pull && docker-compose -f ./docker-compose.yml up -d

stop:
	docker-compose -f ./docker-compose.yml down

.PHONY: docker restart stop
