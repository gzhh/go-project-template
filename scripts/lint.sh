#!/bin/sh

echo "====================="
echo "format source files"
echo "go fmt ./..."
go fmt ./...

echo "====================="
echo "examines Go source code and reports suspicious constructs, detect any suspicious, abnormal, or useless code"
echo "go vet ./..."
echo "====================="
go vet ./...

echo "====================="
echo "golint, start basic lint"
echo "====================="
golint ./...

echo "====================="
echo "golangci-lint, check source files"
echo "====================="
golangci-lint run
