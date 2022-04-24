GO_TEST_OPTIONS = "-race -v"

.PHONY: test
test:
	@go test $(GO_TEST_OPTION) ./...

.PHONY: coverage
coverage:
	@go test $(GO_TEST_OPTION) -coverpkg=./... -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: lint
lint:
	@docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.45.2 golangci-lint run /app/...
