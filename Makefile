.PHONY: build
build: swaggo
	@go build \
		-o ./build/app ./src/

.PHONY: run
run: kill-process build
	./build/app

.PHONY: kill-process
kill-process:
	lsof -i :8081 | awk '$$1 ~ /app/ { print $$2 }' | xargs kill -9 || true

.PHONY: integration
integration:
	@go test -v ./... -cover

.PHONE: swaggo-install
swaggo-install:
	@go get -u github.com/swaggo/swag/cmd/swag

.PHONY: swaggo
swaggo:
	@swag init -g "src/main.go" .