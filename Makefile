test:
	go test -v ./...

format:
	go fmt ./...

lint:
	golangci-lint run
