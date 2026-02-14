.PHONY: test lint fmt

test:
	go test -v -race -coverprofile=coverage.out ./...

lint:
	golangci-lint run

fmt:
	gofmt -s -w .
	goimports -w -local github.com/murasame29/unifi-client-go .

# Release a new version (usage: make release VERSION=v1.0.0)
release:
	@if [ -z "$(VERSION)" ]; then echo "VERSION is required. Usage: make release VERSION=v1.0.0"; exit 1; fi
	git tag $(VERSION)
	git push origin $(VERSION)
