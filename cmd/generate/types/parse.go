// Package types is a package that parses the GDNative headers for type definitions
// to create wrapper structures for Go.
package types

import (
	"fmt"
	"github.com/pinzolo/casee"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Parse will parse the GDNative headers. Takes a list of headers/structs to ignore.
// Definitions in the given headers and definitions
// with the given name will not be added to the returned list of type definitions.
// We'll need to manually create these structures.
func Parse(excludeHeaders, excludeStructs []string) []TypeDef {

	// Get the GOPATH so we can locate the godot headers.
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		panic("$GOPATH is not defined. Run 'export GOPATH=/path/to/go/path' before executing this.")
	}
	packagePath := goPath + "/src/github.com/shadowapex/godot-go"

	// Walk through all of the godot header files
	searchDir := packagePath + "/godot_headers"
	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.Contains(path, ".h") {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// Create a list of all our type definitions
	typeDefinitions := []TypeDef{}

	// Loop through all of the Godot header files and parse the type definitions
	for _, header := range fileList {
		fmt.Println("Parsing header:", header, "...")
		headerName := strings.Replace(header, searchDir+"/", "", 1)

		// Read the header
		content, err := ioutil.ReadFile(header)
		if err != nil {
			panic(err)
		}

		// Find all of the type definitions in the header file
		foundTypesLines := findTypeDefs(content)
		fmt.Println("")

		// After extracting the lines, we can now parse the type definition to
		// a structure that we can use to build a Go wrapper.
		for _, foundTypeLines := range foundTypesLines {
			typeDef := parseTypeDef(foundTypeLines, headerName)

			// Only add the type if it's not in our exclude list.
			if !strInSlice(typeDef.Name, excludeStructs) && !strInSlice(typeDef.HeaderName, excludeHeaders) {
				typeDefinitions = append(typeDefinitions, typeDef)
			}
		}
	}

	return typeDefinitions
}

func parseTypeDef(typeLines []string, headerName string) TypeDef {
	// Create a structure for our type definition.
	typeDef := TypeDef{
		HeaderName: headerName,
		Properties: []TypeDef{},
	}

	// Small function for splitting a line to get the uncommented line and
	// get the comment itself.
	getComment := func(line string) (def, comment string) {
		halves := strings.Split(line, "//")
		def = halves[0]
		if len(halves) > 1 {
			comment = strings.TrimSpace(halves[1])
		}

		return def, comment
	}

	// If the type definition is a single line, handle it a little differently
	if len(typeLines) == 1 {
		// Extract the comment if there is one.
		line, comment := getComment(typeLines[0])

		// Check to see if the property is a pointer type
		if strings.Contains(line, "*") {
			line = strings.Replace(line, "*", "", 1)
			typeDef.IsPointer = true
		}

		// Get the words of the line
		words := strings.Split(line, " ")
		typeDef.Name = strings.Replace(words[len(words)-1], ";", "", 1)
		typeDef.GoName = casee.ToPascalCase(strings.Replace(typeDef.Name, "godot_", "", 1))
		typeDef.Base = words[len(words)-2]
		typeDef.Comment = comment
		typeDef.SimpleType = true

		return typeDef
	}

	// Extract the name of the type.
	lastLine := typeLines[len(typeLines)-1]
	words := strings.Split(lastLine, " ")
	typeDef.Name = strings.Replace(words[len(words)-1], ";", "", 1)

	// Convert the name of the type to a Go name
	typeDef.GoName = casee.ToPascalCase(strings.Replace(typeDef.Name, "godot_", "", 1))

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

		// Extract the comment if there is one.
		line, comment := getComment(line)
		property.Comment = comment

		// Sanitize the line
		line = strings.TrimSpace(line)
		line = strings.Split(line, ";")[0]
		line = strings.Replace(line, "unsigned ", "u", 1)
		line = strings.Replace(line, "const ", "", 1)

		// Split the line by spaces
		words = strings.Split(line, " ")

		// Check to see if the line is just a comment
		if words[0] == "//" {
			continue
		}

		// Set the property details
		if typeDef.Base == "enum" {
			// Strip any commas in the name
			words[0] = strings.Replace(words[0], ",", "", 1)
			property.Name = words[0]
			property.GoName = casee.ToPascalCase(strings.Replace(words[0], "GODOT_", "", 1))
		} else {
			if len(words) < 2 {
				fmt.Println("Skipping irregular line:", line)
				continue
			}
			property.Base = words[0]
			property.Name = words[1]
			property.GoName = casee.ToPascalCase(strings.Replace(words[1], "godot_", "", 1))
		}

		// Check to see if the property is a pointer type
		if strings.Contains(property.Name, "*") {
			property.Name = strings.Replace(property.Name, "*", "", 1)
			property.GoName = strings.Replace(property.GoName, "*", "", 1)
			property.IsPointer = true
		}

		// Append the property to the type definition
		typeDef.Properties = append(typeDef.Properties, property)
	}

	return typeDef
}

// findTypeDefs will return a list of type definition lines.
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

func strInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
