BINARY_NAME=example
DEBUG_NAME=example_debug
MAIN_PACKAGE=./cmd
SRC_DIRS?=$$(go list ./...)

.PHONY: test
test:	export GOCACHE=off
test:
	go test -v $(SRC_DIRS)

.PHONY: build
build:
	go build -v -o $(BINARY_NAME) $(MAIN_PACKAGE)

.PHONY: debug
debug:
	go build -v -gcflags=all="-N -l" -o $(DEBUG_NAME) $(MAIN_PACKAGE)

.PHONY: clean
clean:
	go clean -v $(SRC_DIRS)
	rm -f $(BINARY_NAME) $(DEBUG_NAME)

.DEFAULT_GOAL := debug
