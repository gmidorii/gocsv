.PHONY: linux mac

SRCS    := $(shell find . -type f -name '*.go')

linux: $(SRCS)
  GOOS=linux GOARCH=amd64 go build -o gocsv-linux

mac: $(SRCS)
  GOOS=darwin GOARCH=amd64 go build -o gocsv-mac
