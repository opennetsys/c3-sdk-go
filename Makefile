all: test

.PHONY: test
test:
	@go test -v *.go

.PHONY: deps
deps:
	dep ensure -update -v
