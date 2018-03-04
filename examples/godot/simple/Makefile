all: build run

.PHONY: all

PLATFORM = $(shell uname -s)

build:
ifeq ($(PLATFORM),Linux)
	go build -buildmode=c-shared -o libgodot.so ./src/pong.go
endif 
ifeq ($(PLATFORM),Darwin)
	go build -buildmode=c-shared -o libgodot.dylib ./src/pong.go
endif

run:
	godot Pong.tscn

clean:
ifeq ($(PLATFORM),Linux)
	rm libgodot.so
endif 
ifeq ($(PLATFORM),Darwin)
	rm libgodot.dylib
endif