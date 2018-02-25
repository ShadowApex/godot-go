// Package methods is a package that parses the GDNative headers for type definitions
// of methods
package methods

import (
	"github.com/shadowapex/godot-go/cmd/generate/gdnative"
	"os"
)

// Parse will parse the GDNative headers. Takes a list of headers/structs to ignore.
// Definitions in the given headers and definitions
// with the given name will not be added to the returned list of type definitions.
// We'll need to manually create these structures.
func Parse() gdnative.APIs {

	// Get the GOPATH so we can locate the godot headers.
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		panic("$GOPATH is not defined. Run 'export GOPATH=/path/to/go/path' before executing this.")
	}
	packagePath := goPath + "/src/github.com/shadowapex/godot-go"

	// Parse the GDNative JSON for method data.
	apis := gdnative.Parse(packagePath)
	return apis
}
