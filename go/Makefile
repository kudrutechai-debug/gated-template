.PHONY: check audit lint arch test tools

## The same four the gate runs, in the same order.
check: audit lint arch test

tools:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/fe3dback/go-arch-lint@latest
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.6.1

audit:
	govulncheck ./...

lint:
	golangci-lint run

arch:
	go-arch-lint check

test:
	go test ./... -race -count=1
