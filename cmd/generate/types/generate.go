// Package types is responsible for parsing the Godot headers for type definitions
// and generating Go wrappers around that structure.
package types

import (
	"bytes"
	"github.com/kr/pretty"
	"github.com/pinzolo/casee"
	"github.com/shadowapex/godot-go/cmd/generate/methods"
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
	Headers           []string
	TypeDefinitions   []TypeDef
	MethodDefinitions []Method
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

func (v View) ToGoReturnType(str string) string {
	str = v.ToGoArgType(str, true)
	if strings.Contains(str, "Void") {
		return ""
	}

	return str
}

func (v View) HasReturn(str string) bool {
	if str == "void" || str == "Void" || strings.Contains(str, "void") {
		return false
	}
	return true
}

func (v View) HasPointerReturn(str string) bool {
	if strings.Contains(str, "*") {
		return true
	}
	return false
}

func (v View) IsVoidPointerType(str string) bool {
	switch str {
	case "godot_object *":
		return true
	}
	return false
}

func (v View) IsWcharT(str string) bool {
	if strings.Contains(str, "wchar_t") {
		return true
	}
	return false
}

func (v View) IsDoublePointer(str string) bool {
	if strings.Contains(str, "**") {
		return true
	}
	return false
}

func (v View) ToGoArgType(str string, parseArray bool) string {
	str = strings.Replace(str, "const ", "", -1)
	str = v.ToGoName(str)
	str = strings.Replace(str, "*", "", 1)
	str = strings.TrimSpace(str)

	// If the string still contains a *, it is a list.
	if strings.Contains(str, "*") {
		str = strings.Replace(str, "*", "", 1)
		if parseArray {
			str = "[]" + str
		}
	}

	return str
}

func (v View) ToGoArgName(str string) string {
	if strings.HasPrefix(str, "p_") {
		str = strings.Replace(str, "p_", "", 1)
	}
	if strings.HasPrefix(str, "r_") {
		str = strings.Replace(str, "r_", "", 1)
	}
	str = casee.ToCamelCase(str)

	// Check for any reserved names
	switch str {
	case "type":
		return "aType"
	case "default":
		return "aDefault"
	case "var":
		return "variable"
	case "func":
		return "function"
	case "return":
		return "returns"
	case "interface":
		return "intrfce"
	case "string":
		return "str"
	}

	return str
}

func (v View) IsBasicType(str string) bool {
	switch str {
	case "Uint", "WcharT", "Bool", "Double", "Error", "Int", "Int64T", "Uint64T", "Uint8T", "Uint32T", "Real", "MethodRpcMode", "PropertyHint", "SignedChar", "UnsignedChar", "Vector3Axis":
		return true
	}

	return false
}

// OutputCArg will determine if we need to reference, dereference, etc. an argument
// before passing it to a C function.
func (v View) OutputCArg(arg []string) string {
	argType := arg[0]

	// For basic types, we usually don't pass by pointer.
	if v.IsBasicType(v.ToGoArgType(argType, true)) {
		if v.HasPointerReturn(argType) {
			return "&"
		}
		if argType == "wchar_t" && !v.HasPointerReturn(argType) {
			return "*"
		}
		return ""
	}

	// Non-basic types are returned as pointers. If the C function doesn't want
	// a pointer, we need to dereference the argument.
	if !v.HasPointerReturn(argType) {
		return "*"
	}

	return ""
}

// MethodsList returns all of the methods that match this typedef.
func (v View) MethodsList(typeDef TypeDef) []Method {
	methods := []Method{}

	// Look for all methods that match this typedef name.
	for _, method := range v.MethodDefinitions {
		for _, arg := range method.Arguments {
			argName := arg[1]
			argType := strings.Replace(arg[0], "const", "", 1)
			argType = strings.Replace(argType, "*", "", 1)
			argType = strings.TrimSpace(argType)

			if argType == typeDef.Name && argName == "p_self" {
				methods = append(methods, method)
			}
		}
	}

	return methods
}

func (v View) MethodIsConstructor(method Method) bool {
	if strings.Contains(method.Name, "_new") {
		return true
	}
	return false
}

func (v View) NotSelfArg(str string) bool {
	if str == "self" || str == "p_self" {
		return false
	}
	return true
}

func (v View) ToGoMethodName(typeDef TypeDef, method Method) string {
	methodName := method.Name

	// Replace the typedef in the method name
	methodName = strings.Replace(methodName, typeDef.Name, "", 1)

	// Swap some things around if this is a constructor
	if v.MethodIsConstructor(method) {
		methodName = strings.Replace(methodName, "_new", "", 1)
		methodName = "new_" + typeDef.GoName + "_" + methodName
	}

	if strings.HasPrefix(methodName, "2") {
		methodName = "T" + methodName
	}

	return casee.ToPascalCase(methodName)
}

type Method struct {
	Name       string
	ReturnType string
	Arguments  [][]string
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

	// Parse all available methods
	gdnativeAPI := methods.Parse()

	// Convert the API definitions into a method struct
	allMethodDefinitions := []Method{}
	for _, api := range gdnativeAPI.Core.API {
		method := Method{
			Name:       api.Name,
			ReturnType: api.ReturnType,
			Arguments:  api.Arguments,
		}
		allMethodDefinitions = append(allMethodDefinitions, method)
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
		view.MethodDefinitions = allMethodDefinitions
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

		log.Println("  Running goimports on output:", outFileName+"...")
		GoImports(packagePath + "/gdnative/" + outFileName)
	}

	pretty.Println(allMethodDefinitions)
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

func GoImports(filePath string) {
	cmd := exec.Command("goimports", "-w", filePath)
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		log.Println("Error running goimports:", err)
		panic(stdErr.String())
	}
}
