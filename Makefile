bin::
	[ -d ./dist ] || mkdir -p ./dist
	CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o ./dist/scoper
