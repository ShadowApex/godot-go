#!/bin/bash
PKG_PATH="$GOPATH/src/github.com/shadowapex/godot-go"
TMPLT_PATH="$PKG_PATH/templates"

# Get XML Parser
echo "Ensuring XML parser is available..."
go get github.com/gnewton/chidley
echo "Generating structs based on XML in classes.xml..."
echo "package templates" > $TMPLT_PATH/xmlclass.go
$GOPATH/bin/chidley -e GD -u https://raw.githubusercontent.com/godotengine/godot/master/doc/base/classes.xml >> $TMPLT_PATH/xmlclass.go

# Run our generator
echo "Generating Godot classes..."
go run $GOPATH/src/github.com/shadowapex/godot-go/cmd/generate/generate.go | gofmt > $PKG_PATH/godot/classes.go
