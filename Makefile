.PHONY: pb data lint

PROTOC = protoc
GO_BNDATA = go-bindata
GO=go.exe
pb:
	for f in api/**/*.proto; do \
		$(PROTOC) --go_out=plugins=grpc:. $$f; \
		echo compiled: $$f; \
	done

vet:
	./bin/lint.sh

data:
	$(GO_BINDATA) -o data/bindata.go -pkg data data/*.json

build:
	$(GO) build -i backend/*.go
	$(GO) build -i client/*.go