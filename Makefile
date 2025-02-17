.PHONY: run

.DEFAULT_GOAL := run

SRCS := $(wildcard ./src/*.go)
ifeq ($(OS),Windows_NT)
	OUTPUT := Goverdry.exe
else
	OUTPUT := Goverdry
endif

run: $(SRCS)
	go run $(SRCS)

build: $(SRCS)
	go build -o $(OUTPUT) $(SRCS)

fmt: $(SRCS)
	go fmt $(SRCS)
