package classes

/*
This is meant to parse the godot documentation to generate Go structs of all the godot
objects.
*/

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"github.com/kr/pretty"
	"github.com/pinzolo/casee"
)

// View is a structure that holds the api classes struct, but has additional methods
// attached to it that we can call inside our template.
type View struct {
	API          GDAPI
	APIs         []GDAPI
	Package      string
	PackageMap   map[string]string
	Imports      map[string]map[string]bool
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

func (v View) GetImports() []string {
	imports := []string{}
	for key, _ := range v.Imports[v.API.Name] {
		imports = append(imports, key)
	}

	return imports
}

func (v View) GoName(str string) string {
	return casee.ToPascalCase(strings.Replace(str, "godot_", "", 1))
}

func (v View) GoEmptyReturnType(str string) string {
	if v.IsEnum(str) {
		return "gdnative.NewEmptyInt"
	}
	str = strings.Replace(str, "*", "", 1)
	str = strings.TrimSpace(str)

	if v.IsGodotClass(str) {
		return "gdnative.NewEmptyObject"
	} else {
		return "gdnative.NewEmpty" + v.GoName(str)
	}
}

func (v View) GoNewFromPointerType(str string) string {
	if strings.Contains(str, "enum") || strings.Contains(str, "Enum") {
		return "Int"
	}
	str = strings.Replace(str, "*", "", 1)
	str = strings.TrimSpace(str)
	if v.IsGodotClass(str) {
		return str
	}
	str = casee.ToPascalCase(str)
	return str
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

	// Check to see if the method name is a class name. If so, we need to change
	// it because it will affect calling embedded methods with the same name.
	for _, api := range v.APIs {
		if api.Name == methodString {
			methodString = methodString + "Method"
		}
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
	case "range":
		return "rng"
	case "interface":
		return "intrfce"
	}

	return casee.ToCamelCase(argString)
}

// GoValue will convert the Godot value into a valid Go value.
func (v View) GoValue(returnString string) string {
	// Handle enum values in a specific way
	if strings.Contains(returnString, "enum.") {
		// Strip the enum portion of the string
		returnString = strings.Replace(returnString, "enum.", "", 1)

		// Check to see if this is a class enum or gdnative enum.
		if strings.Contains(returnString, "::") {
			// This is a class enum. Split it to get the class.
			returnSlice := strings.Split(returnString, "::")
			className := returnSlice[0]
			enumName := returnSlice[1]

			// Check for certain this is a class enum
			if v.IsGodotClass(className) {
				return className + enumName
			}
			return "gdnative." + className + enumName
		}

		// This is a gdnative enum.
		return "gdnative." + returnString
	}
	if returnString == "void" {
		return ""
	}
	if v.IsGodotClass(returnString) {
		return returnString
	} else {
		return "gdnative." + casee.ToPascalCase(returnString)
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

func (v View) IsEnum(str string) bool {
	if strings.Contains(str, "enum") || strings.Contains(str, "Enum") {
		return true
	}
	return false
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

func (v View) PackageName(classString string) string {
	str := strings.ToLower(classString)
	switch str {
	case "range":
		return "ranges"
	}
	return str
}

// ResolvePackage will look up the api in the package map for the Go package name
// that we should use.
func (v View) ResolvePackage(curPkg, api string) string {
	if apiPkg, ok := v.PackageMap[api]; ok {
		if curPkg == apiPkg {
			return ""
		}
		return apiPkg
	}
	return "gdnative"
}

// IsGodotClass will determine if the given string is the name of a Godot class
// API. This is used to check if a type is another class or a gdnative base type.
func (v View) IsGodotClass(str string) bool {
	str = strings.Replace(str, "*", "", 1)
	str = strings.TrimSpace(str)
	for api, _ := range v.PackageMap {
		if str == api {
			return true
		}
	}
	return false
}

// HasParentMethod checks to see if the given method exists in any of its parents.
// It will recursively search for any parent with the given method name. This is
// used to generate interfaces for every Godot class type.
func (v View) HasParentMethod(base, method string) bool {
	if base == "" {
		return false
	}

	// Look up the base class
	var baseAPI GDAPI
	for _, api := range v.APIs {
		if api.Name == base {
			baseAPI = api
		}
	}

	if v.HasParentMethod(baseAPI.BaseClass, method) {
		return true
	}

	for _, m := range baseAPI.Methods {
		goMethodName := v.GoMethodName(m.Name)
		if goMethodName == method {
			return true
		}
	}
	return false
}

func Generate() {

	// Get the GOPATH so we can locate our templates.
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		panic("$GOPATH is not defined. Run 'export GOPATH=/path/to/go/path' before executing this.")
	}
	packagePath := goPath + "/src/github.com/shadowapex/godot-go"

	// Get the docs path so we can parse the documentation.
	docsPath := packagePath + "/doc/doc/classes"
	templatePath := packagePath + "/cmd/generate/templates"
	classPath := packagePath + "/godot"

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

	// Prefixes contains a list of package prefixes that will be grouped
	// together if the API starts with this prefix.
	prefixes := []string{
		"animation",
		"audiostream",
		"arvr",
		"audioeffect",
		"bitmap",
		"bullet",
		"dynamicfont",
		"editor",
		"inputevent",
		"ip",
		"physics",
		"popup",
		"reference",
		"regex",
		"resource",
		"scenetree",
		"shader",
		"slider",
		"spatial",
		"streampeer",
		"stylebox",
		"texture",
		"tree",
		"vehicle",
		"videostream",
		"viewport",
		"visibility",
		"visualscript",
		"websocket",
	}

	// Generate a package lookup table for all APIs.
	view.PackageMap = map[string]string{}
	for _, api := range view.APIs {
		packageName := view.PackageName(api.Name)
		for _, prefix := range prefixes {
			if strings.HasPrefix(packageName, prefix) {
				packageName = prefix
			}
		}
		view.PackageMap[api.Name] = packageName
	}

	// Find all of the imports for each API
	view.Imports = map[string]map[string]bool{}
	for _, api := range view.APIs {
		view.Imports[api.Name] = map[string]bool{}
		for _, method := range api.Methods {
			for _, arg := range method.Arguments {
				if view.IsGodotClass(arg.Type) {
					pkg := view.PackageMap[arg.Type]
					view.Imports[api.Name][pkg] = true
				}
			}
			if view.IsGodotClass(method.ReturnType) {
				pkg := view.PackageMap[method.ReturnType]
				view.Imports[api.Name][pkg] = true
			}
		}
	}
	pretty.Println(view.Imports)

	// Loop through all of the APIs and generate packages for them.
	for _, api := range view.APIs {
		// Skip classes to generate.
		if !view.IsValidClass(api.Name, api.BaseClass) {
			continue
		}

		// Get the package name to generate
		packageName := view.PackageMap[api.Name]
		outFileName := view.PackageName(api.Name) + ".gen.go"

		// Set the current API
		view.API = api
		view.Package = packageName

		// Write the file using our template.
		WriteTemplate(
			packagePath+"/cmd/generate/templates/class.go.tmpl",
			classPath+"/"+outFileName,
			view,
		)

		// Run gofmt on the generated Go file.
		log.Println("  Running gofmt on output:", outFileName+"...")
		GoFmt(classPath + "/" + outFileName)

		log.Println("  Running goimports on output:", outFileName+"...")
		GoImports(classPath + "/" + outFileName)

	}

	// Generate the conversion function to convert based on class name.
	log.Println("Generating conversion functions.")
	outFileName := "convert.gen.go"
	WriteTemplate(
		packagePath+"/cmd/generate/templates/convert.go.tmpl",
		classPath+"/"+outFileName,
		view,
	)

	// Run gofmt and goimports on the conversions
	log.Println("  Running gofmt on output:", outFileName+"...")
	GoFmt(classPath + "/" + outFileName)

	log.Println("  Running goimports on output:", outFileName+"...")
	GoImports(classPath + "/" + outFileName)

	log.Println(len(view.APIs))
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
