// Package types is responsible for parsing the Godot headers for type definitions
// and generating Go wrappers around that structure.
package types

import (
	"log"
	"os"
	"sort"
	"text/template"
)

// View is a structure that holds the api struct, so it can be used inside
// our temaplte.
type View struct {
	Headers         []string
	TypeDefinitions []TypeDef
}

// Generate will generate Go wrappers for all Godot base types
func Generate() {

	// Get the GOPATH so we can locate our templates.
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		panic("$GOPATH is not defined. Run 'export GOPATH=/path/to/go/path' before executing this.")
	}
	packagePath := goPath + "/src/github.com/shadowapex/godot-go"

	// Create a structure for our template view. This will contain all of
	// the data we need to construct our Go wrappers.
	var view View

	// Parse the Godot header files for type definitions
	view.TypeDefinitions = Parse()
	view.Headers = []string{}

	// Collect all of the headers we need to use in our template.
	headers := map[string]bool{}
	for _, typeDef := range view.TypeDefinitions {
		headers[typeDef.HeaderName] = true
	}
	for header := range headers {
		view.Headers = append(view.Headers, header)
	}
	sort.Strings(view.Headers)

	// Loop through all of the type definitions we have.
	log.Println("Generating Go wrappers for Godot base types...")
	WriteTemplate(
		packagePath+"/cmd/generate/templates/types.go.tmpl",
		packagePath+"/gdnative/types_tmp.go",
		view,
	)

}

func WriteTemplate(templatePath, outputPath string, view View) {
	// Create a template from our template file.
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	// Open the output file for writing
	f, err := os.Create(outputPath)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	// Write the template with the given view.
	err = t.Execute(f, view)
	if err != nil {
		panic(err)
	}
}
