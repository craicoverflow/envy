test:
	go test -v ./...
.PHONY: test

test/coverage:
	go test -coverprofile=coverage.out
.PHONY: test/coverage

coverage:
	go tool cover -html=coverage.out
.PHONY: coverage

format:
	go fmt ./...

lint:
	golangci-lint run
