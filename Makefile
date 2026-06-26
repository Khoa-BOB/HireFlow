run:
	go run ./cmd/server/main.go

fmt:
	go fmt ./...

lint:
	golangci-lint run

docs:
	echo "OpenAPI docs are in /internal/docs/openapi.yaml"