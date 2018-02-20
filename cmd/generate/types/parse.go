// Package types is a package that parses the GDNative headers for type definitions
// to create wrapper structures for Go.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type TypeDef struct {
	Name       string
	Base       string
	Properties []TypeDef
}

func main() {
	// Get the GOPATH so we can locate the godot headers.
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		panic("$GOPATH is not defined. Run 'export GOPATH=/path/to/go/path' before executing this.")
	}
	packagePath := goPath + "/src/github.com/shadowapex/godot-go"

	fmt.Println("vim-go")
	// Read the header
	content, err := ioutil.ReadFile(packagePath + "/godot_headers/gdnative/gdnative.h")
	if err != nil {
		panic(err)
	}

	// Find all of the type definitions in the header file
	foundTypes := findTypeDefs(content)
	fmt.Println("")

	// After extracting the lines, we can now parse the type definition to
	// a structure that we can use to build a Go wrapper.
	for _, foundType := range foundTypes {
		typeDef := parseTypeDef(foundType)
		fmt.Println(typeDef)
	}
}

func parseTypeDef(typeLines []string) TypeDef {
	// Create a structure for our type definition.
	typeDef := TypeDef{
		Properties: []TypeDef{},
	}

	// If the type definition is a single line, handle it a little differently
	if len(typeLines) == 1 {
		// Get the words of the line
		words := strings.Split(typeLines[0], " ")
		typeDef.Name = strings.Replace(words[len(words)-1], ";", "", 1)
		typeDef.Base = words[len(words)-2]

		return typeDef
	}

	// Extract the name of the type.
	lastLine := typeLines[len(typeLines)-1]
	words := strings.Split(lastLine, " ")
	typeDef.Name = strings.Replace(words[len(words)-1], ";", "", 1)

	// Extract the base type
	firstLine := typeLines[0]
	words = strings.Split(firstLine, " ")
	typeDef.Base = words[1]

	// Extract the properties from the type
	properties := typeLines[1 : len(typeLines)-1]

	// Loop through each property line
	for _, line := range properties {
		// Skip function definitions
		if strings.Contains(line, ")") {
			continue
		}

		// Create a type definition for the property
		property := TypeDef{}

		// Sanitize the line
		line = strings.TrimSpace(line)
		line = strings.Split(line, ";")[0]
		line = strings.Replace(line, "unsigned ", "u", 1)
		line = strings.Replace(line, "const ", "", 1)

		// Split the line by spaces
		words = strings.Split(line, " ")

		// Set the property details
		if typeDef.Base != "enum" {
			property.Base = words[0]
			property.Name = words[1]
		} else {
			property.Name = words[0]
		}

		// Append the property to the type definition
		typeDef.Properties = append(typeDef.Properties, property)
	}

	return typeDef
}

func findTypeDefs(content []byte) [][]string {
	lines := strings.Split(string(content), "\n")

	// Create a structure that will hold the lines that define the
	// type.
	singleType := []string{}
	foundTypes := [][]string{}
	var typeFound bool = false
	for i, line := range lines {

		// Search the line for type definitions
		if strings.Contains(line, "typedef ") {
			fmt.Println("Line", i, ": Type found on line:", line)
			typeFound = true

			// Check to see if this is a single line type. If it is,
			// we're done.
			if strings.Contains(line, ";") {
				// Skip if this is a function definition
				if strings.Contains(line, ")") {
					typeFound = false
					continue
				}

				fmt.Println("Line", i, ": Short type definition found.")
				singleType = append(singleType, line)
				typeFound = false
				foundTypes = append(foundTypes, singleType)
				singleType = []string{}
			}
		}

		// If a type was found, keep appending our struct lines until we
		// reach the end of the definition.
		if typeFound {
			fmt.Println("Line", i, ": Appending line for type found:", line)

			// Keep adding the lines to our list of lines until we
			// reach an end bracket.
			singleType = append(singleType, line)

			if strings.Contains(line, "}") {
				fmt.Println("Line", i, ": Found end of type definition.")
				typeFound = false
				foundTypes = append(foundTypes, singleType)
				singleType = []string{}
			}
		}
	}

	return foundTypes
}
