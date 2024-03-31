build:
	@go build -o bin/go-dict

run: build
	./bin/go-dict

test:
	go test -v ./...
