BINARY_NAME := get-secret-value

all: vet build

build: FORCE
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(CURDIR)/bin/${BINARY_NAME}_linux_amd64
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o $(CURDIR)/bin/${BINARY_NAME}_darwin_amd64

vet:
	go vet ./...
clean:
	$(RM) -r $(CURDIR)/bin

FORCE:

.PHONY: all build vet clean
