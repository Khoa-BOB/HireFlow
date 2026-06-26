.PHONY: run fmt lint docs
run:
	go run ./cmd/server/main.go

fmt:
	go fmt ./...

lint:
	golangci-lint run

docs:
	echo "OpenAPI docs are in /internal/docs/openapi.yaml"

migration:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	export $$(grep -v '^#' .env | xargs) && \
	migrate -path migrations -database "$$DATABASE_URL" up

migrate-down:
	export $$(grep -v '^#' .env | xargs) && \
	migrate -path migrations -database "$$DATABASE_URL" down 1

migrate-force:
	export $$(grep -v '^#' .env | xargs) && \
	migrate -path migrations -database "$$DATABASE_URL" force $(version)