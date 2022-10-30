.PHONY: config
config:
	$(info [Create build config])
	@go mod tidy

###############################################################################
# static checks
###############################################################################

.PHONY: fmt
fmt:
	gofmt -l -s -w .

.PHONY: install-static-check-tools
install-static-check-tools:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- -b $(GOPATH)/bin v1.48.0
	@go install github.com/matrixorigin/linter/cmd/molint@latest
	@go install github.com/apache/skywalking-eyes/cmd/license-eye@v0.4.0


.PHONY: static-check
static-check: config
	$(CGO_OPTS) go vet -vettool=`which molint` ./...
	$(CGO_OPTS) license-eye -c .licenserc.yml header check
	$(CGO_OPTS) license-eye -c .licenserc.yml dep check
	$(CGO_OPTS) golangci-lint run -c .golangci.yml ./...
