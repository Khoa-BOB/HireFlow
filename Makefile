run:
	go run ./cmd/server/main.go

fmt:
	go fmt ./...

lint:
	golangci-lint run