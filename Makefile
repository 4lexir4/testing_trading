build:
	go build -o bin/testing_trading

run: build
	./bin/testing_trading

test:
	go test -v ./...
