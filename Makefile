.PHONY: pb data lint

PROTOC = protoc
GO_BNDATA = go-bindata

pb:
	for f in pb/**/*.proto; do \
		$(PROTOC) --go_out=plugins=grpc:. $$f; \
		echo compiled: $$f; \
	done

vet:
	./bin/lint.sh

data:
	$(GO_BINDATA) -o data/bindata.go -pkg data data/*.json

