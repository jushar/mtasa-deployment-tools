#!/bin/bash

# Build go project
echo "Building Worker Server..."
go build -o ./workerserver *.go
