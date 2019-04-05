LINT_TOOLS=\
	golang.org/x/lint/golint \
	golang.org/x/tools/cmd/goimports \
	github.com/kisielk/errcheck \
	honnef.co/go/tools/cmd/staticcheck

LINT_PACKAGES = $(shell go list ./...)
FORMAT_PACKAGES = $(foreach pkg,$(LINT_PACKAGES),$(shell go env GOPATH)/src/$(pkg))

.PHONY: bootstrap-lint-tools
bootstrap-lint-tools:
	@for tool in $(LINT_TOOLS) ; do \
		echo "Installing/Updating $$tool" ; \
		GO111MODULE=off go get -u $$tool; \
	done

.PHONY: lint
lint: fmt vet staticcheck errcheck

.PHONY: fmt
fmt:
	goimports -l $(FORMAT_PACKAGES) | grep -E '.'; test $$? -eq 1
	gofmt -l $(FORMAT_PACKAGES) | grep -E '.'; test $$? -eq 1

.PHONY: vet
vet:
	go vet -v $(LINT_PACKAGES)

.PHONY: staticcheck
staticcheck:
	staticcheck $(LINT_PACKAGES)

.PHONY: errcheck
errcheck:
	errcheck -ignore 'fmt:[FS]?[Pp]rint*' -exclude .errcheckignore $(LINT_PACKAGES)
