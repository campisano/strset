SOURCE_FOLDERS?=$$(go list ./... | grep -v /vendor/)
MAIN_PACKAGE=genset/cmd
GDB_TARGET=example

.PHONY: test
test:	export GOCACHE=off
test:
	go test $(SOURCE_FOLDERS)

.PHONY: clean
clean:
	go clean $(SOURCE_FOLDERS)
	rm -f $(GDB_TARGET)

.PHONY: debug
debug:
	go build -v -gcflags=all="-N -l" -o $(GDB_TARGET) $(MAIN_PACKAGE)

.DEFAULT_GOAL := debug
