# Makefile

Target=sns-agent

Go=go
Package=main
Version=$(shell cat VERSION)
GitCommit=$(shell git rev-parse HEAD)
BuildTime=$(shell date '+%Y-%m-%d_%H:%M:%S')

.PHONY: build linux test clean

all: build

build:
	$(Go) build -o $(Target) -ldflags "-X $(Package).Version=$(Version) -X $(Package).GitCommit=$(GitCommit) -X $(Package).BuildTime=$(BuildTime)"

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(Go) build -o $(Target) -ldflags "-X $(Package).Version=$(Version) -X $(Package).GitCommit=$(GitCommit) -X $(Package).BuildTime=$(BuildTime)"

test:
	$(Go) test ./...

clean:
	rm -f $(Target)

