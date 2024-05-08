.PHONY: install

#binary name
BINARY_NAME=dict

#installation dir
INSTALL_DIR=/usr/local/bin/

all: build

build:
	@go build -o bin/go-dict

run: build
	./bin/go-dict

test:
	go test -v ./...

install: build
	sudo cp $(BINARY_NAME) $(INSTALL_DIR)
