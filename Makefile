all: lint test build

lint:
	echo "go fmt"
	go fmt ./...
	echo "go vet"
	go vet ./...
	echo "go lint"
	# golint ./...
	echo "golangci lint"
	# golangci-lint run

test:
	go test -test.bench=".*" -count=1 -v ./...

build:
	go build -o demo -v -a -ldflags '-w -s' -tags=jsoniter cmd/main.go