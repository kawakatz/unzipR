# Reference: https://github.com/projectdiscovery/httpx/blob/master/Makefile
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY=bin/unzipR
GOBIN=~/go/bin
    
all: build
build:
		$(GOBUILD) -v -o "bin/unzipR" cmd/unzipR/unzipR.go
test: 
		$(GOTEST) -v ./...
tidy:
		$(GOMOD) tidy
install:	$(BINARY)
		install -s $(BINARY) $(GOBIN)