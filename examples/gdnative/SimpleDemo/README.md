# GDNative Simple Example

This example demonstrates how you can use the `gdnative` Go package to create
your own Godot bindings with Go.

## Compiling

Dependencies:
 * golang 1.6+

### CGO (Cross platform)
You can use golang to compile the library if you have it installed:

    go build -v -buildmode=c-shared -o libsimple.so ./src/simple.go && mv libsimple.so ./bin

## Usage

Create a new object using `load("res://SIMPLE.gdns").new()`

This object has following methods you can use:
 * `get_data()`

