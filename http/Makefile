.PHONY: run
run: imports
	go run .

.PHONY: debug
debug: imports
	dlv debug .

.PHONY: format
format:
	gofmt -w .

.PHONY: imports
imports: format
	goimports -w .

.PHONY: dep
dep:
	dep ensure
