#!/bin/bash

GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o /app/scoper
