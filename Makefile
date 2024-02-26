.PHONY: help install build 

help:
	@echo "test - run the program"
	@echo "install - install the program"
	@echo "build - build the program"
	@echo "test-build - build the program and run the test"

install:
	go install cmd/gen/main.go

build:
	go build -o bin/garch cmd/gen/main.go

test-build:
	./bin/garch gen -n Pixle

app-test:
	go run cmd/gen/main.go gen -n Pixle -u meh

app-help:
	go run cmd/gen/main.go --help