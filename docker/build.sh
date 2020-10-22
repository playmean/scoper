#!/bin/bash

CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o /app/scoper
