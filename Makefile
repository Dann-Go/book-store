BINARY_NAME = book-store

build:
	go build -o ${BINARY_NAME} ./cmd/api

test:
	go test ./...

