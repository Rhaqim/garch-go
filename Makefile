.PHONY: test install build

test:
	go run cmd/gen/main.go gen --title Pixle

install:
	go install cmd/gen/main.go

build:
	go build -o bin/garch cmd/gen/main.go