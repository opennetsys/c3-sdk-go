all: test

.PHONY: deps
deps:
	@echo "running dep ensure..." && \
	dep ensure -v && \
	$(MAKE) gxundo

.PHONY: gxundo
gxundo:
	@bash scripts/gxundo.sh vendor/

.PHONY: test
test:
	@go test -v *.go

.PHONY: deps
deps:
	dep ensure -update -v  && \
	$(MAKE) gxundo

.PHONY: gxundo
gxundo:
	@bash scripts/gxundo.sh vendor/

.PHONY: install/gxundo
install/gxundo:
	@mkdir -p scripts && \
	wget https://raw.githubusercontent.com/c3systems/gxundo/master/gxundo.sh \
	-O scripts/gxundo.sh && \
	chmod +x scripts/gxundo.sh

