// Package types is responsible for parsing the Godot headers for type definitions
// and generating Go wrappers around that structure.
package types

import (
	"bytes"
	"github.com/kr/pretty"
	"github.com/pinzolo/casee"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"
)

// View is a structure that holds the api struct, so it can be used inside
// our temaplte.
type View struct {
	Headers         []string
	TypeDefinitions []TypeDef
}

// IsValidProperty will determine if we should be generating the given property
// in our Go structure.
func (v View) IsValidProperty(prop TypeDef) bool {
	if strings.Contains(prop.Name, "_touch_that") {
		return false
	}
	return true
}

// IsGodotBaseType will check to see if the given simple type definition is defining
// a built-in C type or a Godot type.
func (v View) IsGodotBaseType(typeDef TypeDef) bool {
	if strings.Contains(typeDef.Base, "godot_") {
		return true
	}
	return false
}

// ToGoBaseType will convert a base type name to the correct Go base type.
func (v View) ToGoBaseType(base string) string {
	switch base {
	case "float":
		return "float64"
	case "wchar_t":
		return "string"
	}

	return base
}

func (v View) ToGoName(str string) string {
	str = strings.Replace(str, "godot_", "", 1)
	str = strings.Replace(str, "GODOT_", "", 1)
	return casee.ToPascalCase(str)
}

// Generate will generate Go wrappers for all Godot base types
func Generate() {

	// Get the GOPATH so we can locate our templates.
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		panic("$GOPATH is not defined. Run 'export GOPATH=/path/to/go/path' before executing this.")
	}
	packagePath := goPath + "/src/github.com/shadowapex/godot-go"

	// Set up headers/structs to ignore. Definitions in the given headers
	// with the given name will not be added to the returned list of type definitions.
	// We'll need to manually create these structures.
	ignoreHeaders := []string{
		"pluginscript/godot_pluginscript.h",
	}
	ignoreStructs := []string{
		"godot_char_type",
		"godot_gdnative_api_struct",
		"godot_gdnative_core_api_struct",
		"godot_gdnative_ext_arvr_api_struct",
		"godot_gdnative_ext_nativescript_1_1_api_struct",
		"godot_gdnative_ext_nativescript_api_struct",
		"godot_gdnative_ext_pluginscript_api_struct",
		"godot_gdnative_init_options",
		"godot_instance_binding_functions",
		"godot_instance_create_func",
		"godot_instance_destroy_func",
		"godot_instance_method",
		"godot_method_attributes",
		"godot_property_get_func",
		"godot_property_set_func",
		"godot_property_usage_flags",
	}

	// Parse the Godot header files for type definitions
	allTypeDefinitions := Parse(ignoreHeaders, ignoreStructs)

	// Create a map of the type definitions by header name
	defMap := map[string][]TypeDef{}

	// Organize the type definitions by header name
	for _, typeDef := range allTypeDefinitions {
		if _, ok := defMap[typeDef.HeaderName]; ok {
			defMap[typeDef.HeaderName] = append(defMap[typeDef.HeaderName], typeDef)
		} else {
			defMap[typeDef.HeaderName] = []TypeDef{typeDef}
		}
	}
	pretty.Println(defMap)

	// Loop through each header name and generate the Go code in a file based
	// on the header name.
	log.Println("Generating Go wrappers for Godot base types...")
	for headerName, typeDefs := range defMap {
		// Convert the header name into the Go filename
		headerPath := strings.Split(headerName, "/")
		outFileName := strings.Replace(headerPath[len(headerPath)-1], ".h", ".gen.go", 1)
		outFileName = strings.Replace(outFileName, "godot_", "", 1)

		log.Println("  Generating Go code for:", outFileName+"...")

		// Create a structure for our template view. This will contain all of
		// the data we need to construct our Go wrappers.
		var view View

		// Add the type definitions for this file to our view.
		view.TypeDefinitions = typeDefs
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

		// Write the file using our template.
		WriteTemplate(
			packagePath+"/cmd/generate/templates/types.go.tmpl",
			packagePath+"/gdnative/"+outFileName,
			view,
		)

		// Run gofmt on the generated Go file.
		log.Println("  Running gofmt on output:", outFileName+"...")
		GoFmt(packagePath + "/gdnative/" + outFileName)
	}
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

func GoFmt(filePath string) {
	cmd := exec.Command("gofmt", "-w", filePath)
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		log.Println("Error running gofmt:", err)
		panic(stdErr.String())
	}
}
