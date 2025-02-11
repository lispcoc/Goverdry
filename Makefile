.PHONY: run

.DEFAULT_GOAL := run

SRCS := $(wildcard ./*.go)

run: 
	go run $(SRCS)

build:
	go build $(SRCS)
