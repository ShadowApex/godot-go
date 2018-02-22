#!/bin/bash
export PKG_PATH="$GOPATH/src/github.com/shadowapex/godot-go"

# Run code generation
go run $PKG_PATH/cmd/generate/main.go
