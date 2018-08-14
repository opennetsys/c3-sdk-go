all: build

.PHONY: build
build: clean
	@go build -o c3_sdk.so -buildmode=c-shared "${GOPATH}/src/github.com/c3systems/c3-sdk-go/lib/c/c.go"

.PHONY: clean
clean:
	@-rm *.so 2>/dev/null || true
