.PHONY: run

.DEFAULT_GOAL := run

SRCS := $(wildcard ./*.go)

run: $(SRCS)
	go run $(SRCS)

build: $(SRCS)
	go build $(SRCS)
