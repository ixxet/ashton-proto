BUF ?= buf
GO ?= go

.PHONY: generate lint test check breaking clean

generate:
	$(BUF) generate

lint:
	$(BUF) lint

test:
	$(GO) test ./...

check: lint generate test

breaking:
	$(BUF) breaking --against '.git#branch=main'

clean:
	rm -rf gen
