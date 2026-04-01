BUF ?= buf

.PHONY: generate lint breaking clean

generate:
	$(BUF) generate

lint:
	$(BUF) lint

breaking:
	$(BUF) breaking --against '.git#branch=main'

clean:
	rm -rf gen
