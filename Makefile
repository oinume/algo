LINT_TOOLS=\
	golang.org/x/lint/golint \
	golang.org/x/tools/cmd/goimports \
	github.com/kisielk/errcheck \
	honnef.co/go/tools/cmd/staticcheck

GO_TEST_OPTIONS = "-race -v"
LINT_PACKAGES = $(shell go list ./...)
FORMAT_PACKAGES = $(foreach pkg,$(LINT_PACKAGES),$(shell go env GOPATH)/src/$(pkg))

.PHONY: bootstrap-lint-tool
bootstrap-lint-tool:
	@cd tool && for tool in $(LINT_TOOLS) ; do \
		echo "Installing/Updating $$tool" ; \
		GO111MODULE=on go install $$tool; \
	done

.PHONY: test
test:
	go test $(GO_TEST_OPTION) ./...

.PHONY: coverage
coverage:
	go test $(GO_TEST_OPTION) -coverpkg=./... -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: lint
lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.45.2 golangci-lint run /app/...
