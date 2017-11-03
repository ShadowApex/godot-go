package main

/*
This is meant to parse the godot documentation to generate Go structs of all the godot
objects.
*/

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"github.com/pinzolo/casee"
)

// GDAPI is a structure for parsed JSON from godot_api.json.
type GDAPI struct {
	APIType      string           `json:"api_type"`
	BaseClass    string           `json:"base_class"`
	Constants    map[string]int64 `json:"constants"`
	Enums        []GDEnums        `json:"enums"`
	Methods      []GDMethod       `json:"methods"`
	Name         string           `json:"name"`
	Properties   []GDProperty     `json:"properties"`
	Signals      []GDSignal       `json:"signals"`
	Singleton    bool             `json:"singleton"`
	Instanciable bool             `json:"instanciable"`
	IsReference  bool             `json:"is_reference"`
}

// ByName is used for sorting GDAPI objects by name
type ByName []GDAPI

func (c ByName) Len() int           { return len(c) }
func (c ByName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByName) Less(i, j int) bool { return c[i].Name < c[j].Name }

type GDEnums struct {
	Name   string           `json:"name"`
	Values map[string]int64 `json:"values"`
}

// ByEnumName is used for sorting GDAPI objects by name
type ByEnumName []GDEnums

func (c ByEnumName) Len() int           { return len(c) }
func (c ByEnumName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByEnumName) Less(i, j int) bool { return c[i].Name < c[j].Name }

type GDMethod struct {
	Arguments    []GDArgument `json:"arguments"`
	HasVarargs   bool         `json:"has_varargs"`
	IsConst      bool         `json:"is_const"`
	IsEditor     bool         `json:"is_editor"`
	IsFromScript bool         `json:"is_from_script"`
	IsNoscript   bool         `json:"is_noscript"`
	IsReverse    bool         `json:"is_reverse"`
	IsVirtual    bool         `json:"is_virtual"`
	Name         string       `json:"name"`
	ReturnType   string       `json:"return_type"`
}

// ByMethodName is used for sorting GDAPI objects by name
type ByMethodName []GDMethod

func (c ByMethodName) Len() int           { return len(c) }
func (c ByMethodName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByMethodName) Less(i, j int) bool { return c[i].Name < c[j].Name }

type GDArgument struct {
	DefaultValue    string `json:"default_value"`
	HasDefaultValue bool   `json:"has_default_value"`
	Name            string `json:"name"`
	Type            string `json:"type"`
}

type GDProperty struct {
	Getter string `json:"getter"`
	Name   string `json:"name"`
	Setter string `json:"setter"`
	Type   string `json:"type"`
}

// ByPropertyName is used for sorting GDAPI objects by name
type ByPropertyName []GDProperty

func (c ByPropertyName) Len() int           { return len(c) }
func (c ByPropertyName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByPropertyName) Less(i, j int) bool { return c[i].Name < c[j].Name }

type GDSignal struct {
	Arguments []GDArgument `json:"arguments"`
	Name      string       `json:"name"`
}

// BySignalName is used for sorting GDAPI objects by name
type BySignalName []GDSignal

func (c BySignalName) Len() int           { return len(c) }
func (c BySignalName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c BySignalName) Less(i, j int) bool { return c[i].Name < c[j].Name }

// GDAPIDoc is a structure for parsed documentation.
type GDAPIDoc struct {
	Name        string        `xml:"name,attr"`
	Description string        `xml:"description"`
	Methods     []GDMethodDoc `xml:"methods>method"`
}

type GDMethodDoc struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description"`
}

// View is a structure that holds the api classes struct, but has additional methods
// attached to it that we can call inside our template.
type View struct {
	APIs         []GDAPI
	Header       string
	ClassDocs    map[string]string
	MethodDocs   map[string]map[string]string
	SingletonMap map[string]bool
}

// ClassDoc returns the class documentation for the given class.
func (v View) ClassDoc(class string) string {
	class = v.GoClassName(class)
	if _, ok := v.ClassDocs[class]; ok {
		return v.UltraTrim(v.ClassDocs[class])
	}
	return "Undocumented"
}

// MethodDoc returns the method documentation for a given class method.
func (v View) MethodDoc(class, method string) string {
	if _, ok := v.MethodDocs[class][method]; ok {
		return v.UltraTrim(v.MethodDocs[class][method])
	}
	return "Undocumented"
}

// Trim can be used inside the template to trim starting and ending whitespace.
func (v View) Trim(str string) string {
	return strings.TrimSpace(str)
}

// UltraTrim will use RegEx to clean all extra whitespace.
func (v View) UltraTrim(input string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := re_leadclose_whtsp.ReplaceAllString(input, "")
	final = re_inside_whtsp.ReplaceAllString(final, " ")

	return final
}

// GoClassName will convert any _<Name> classes into normal CamelCase names.
func (v View) GoClassName(classString string) string {
	if strings.HasPrefix(classString, "_") {
		return strings.Replace(classString, "_", "", 1)
	}
	return classString
}

// GoMethodName will convert the snake_case'd version of the Godot method name
// into a CamelCase version that is the Go convention.
func (v View) GoMethodName(methodString string) string {
	isPrivate := false

	// Convert the snake_case name to CamelCase
	if strings.HasPrefix(methodString, "_") {
		methodString = strings.Replace(methodString, "_", "X_", 1)
		isPrivate = true
	}
	methodString = casee.ToPascalCase(methodString)

	if isPrivate {
		return strings.Replace(methodString, "X", "X_", 1)
	}
	return methodString
}

// GoArgName will check for Go reserved keywords like "type" when used as argument
// names and convert them, so we don't get compile errors.
func (v View) GoArgName(argString string) string {
	switch argString {
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
	}

	return casee.ToCamelCase(argString)
}

// GoValue will convert the Godot value into a valid Go value.
func (v View) GoValue(returnString string) string {
	// TODO: Right now we're converting any enum types to int64. We should
	// look into creating types for each of these maybe? LOL
	if strings.Contains(returnString, "enum.") {
		returnString = "int"
	}
	switch returnString {
	case "String":
		return "string"
	case "int":
		return "int64"
	case "uint":
		return "uint64"
	case "float":
		return "float64"
	case "bool":
		return "bool"
	case "void":
		return ""
	default:
		return "*" + returnString
	}
}

// IsValidClass will check the class to see if we should generate Go bindings for
// it.
func (v View) IsValidClass(classString, inheritsString string) bool {
	runeString := []rune(classString)
	if strings.HasPrefix(classString, "@") {
		return false
	}
	if strings.HasPrefix(classString, "_") {
		return false
	}
	if unicode.IsLower(runeString[0]) {
		return false
	}
	// NOTE: Object is the only class without inheritance that we should generate.
	if classString == "Object" {
		return true
	}
	// NOTE: We should only autogenerate classes that inherit from "Object".
	if v.Trim(inheritsString) == "" {
		return false
	}

	return true
}

func (v View) SetClassName(classString string, singleton bool) string {
	if singleton {
		return casee.ToCamelCase(classString)
	}
	return classString
}

func (v View) SetBaseClassName(baseClass string) string {
	return v.SetClassName(baseClass, v.SingletonMap[baseClass])
}
func main() {

	// Get the basepath so we know where to look for our JSON API and template file.
	templatePath := os.Getenv("TMPLT_PATH")
	if templatePath == "" {
		panic("$TMPLT_PATH environment variable was not set. Be sure to run this from generate.sh!")
	}
	docsPath := os.Getenv("GODOT_DOC_PATH")
	if docsPath == "" {
		panic("$GODOT_DOC_PATH environment variable was not set. Be sure to run this from generate.sh!")
	}

	// Get our documentation that was pulled down from generate.sh.
	docFiles, err := ioutil.ReadDir(docsPath)
	if err != nil {
		panic(err)
	}

	// Loop through all of the documentation files and parse them. We will use this for
	// populating the comment strings in classes.go.
	classDocs := map[string]string{}
	methodDocs := map[string]map[string]string{}
	for _, docFile := range docFiles {
		body, err := ioutil.ReadFile(docsPath + "/" + docFile.Name())
		if err != nil {
			panic(err)
		}
		var obj GDAPIDoc
		xml.Unmarshal(body, &obj)

		// Populate our class docs
		classDocs[obj.Name] = obj.Description
		methodDocs[obj.Name] = map[string]string{}

		// Populate our method docs
		for _, method := range obj.Methods {
			methodDocs[obj.Name][method.Name] = method.Description
		}
	}

	// Open our godot_api.json file
	body, err := ioutil.ReadFile(templatePath + "/godot_api.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal the JSON into a defined structure.
	var view View
	json.Unmarshal(body, &view.APIs)
	view.Header = `
//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/classes.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------
`
	// Add our documentation to our view.
	view.ClassDocs = classDocs
	view.MethodDocs = methodDocs

	// Store all objects singleton value so it can be looked up later.
	view.SingletonMap = map[string]bool{}
	for _, api := range view.APIs {
		view.SingletonMap[api.Name] = api.Singleton
	}

	// Sort the APIs so they will be generated in order.
	sort.Sort(ByName(view.APIs))
	for _, api := range view.APIs {
		sort.Sort(ByMethodName(api.Methods))
		sort.Sort(ByEnumName(api.Enums))
		sort.Sort(ByPropertyName(api.Properties))
		sort.Sort(BySignalName(api.Signals))
	}

	// List out template file
	templateFile := templatePath + "/classes.go.template"

	// Create a template from our template file.
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}
	templateBuffer := bytes.NewBufferString("")
	err = t.Execute(templateBuffer, view)
	if err != nil {
		log.Fatal("Unable to apply template:", err)
	}

	// Output our template to STDOUT
	fmt.Println(templateBuffer.String())

}
