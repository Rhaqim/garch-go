.PHONY: test install build

test:
	go run cmd/gen/main.go gen -n Pixle -u meh

install:
	go install cmd/gen/main.go

build:
	go build -o bin/garch cmd/gen/main.go

test-build:
	./bin/garch gen -n Pixle