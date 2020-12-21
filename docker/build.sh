#!/bin/bash

GOOS=linux go build -a -ldflags '-s -w -extldflags "-static"' -o /app/scoper
