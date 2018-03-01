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

	"github.com/pinzolo/casee"
)

// View is a structure that holds the api classes struct, but has additional methods
// attached to it that we can call inside our template.
type View struct {
	API          GDAPI
	APIs         []GDAPI
	Package      string
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
	case "range":
		return "rng"
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
		return "gdnative.String"
	case "int":
		return "gdnative.Int"
	case "uint":
		return "gdnative.Uint"
	case "float":
		return "gdnative.Float"
	case "bool":
		return "gdnative.Bool"
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

func (v View) PackageName(classString string) string {
	str := strings.ToLower(classString)
	switch str {
	case "range":
		return "ranges"
	}
	return str
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
	classPath := packagePath + "/godot/class"

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

	// Loop through all of the APIs and generate packages for them.
	for _, api := range view.APIs {
		// Skip classes to generate.
		if !view.IsValidClass(api.Name, api.BaseClass) {
			continue
		}

		// Get the package name to generate
		packageName := view.PackageName(api.Name)
		outFileName := packageName + ".go"
		for _, prefix := range prefixes {
			if strings.HasPrefix(packageName, prefix) {
				packageName = prefix
			}
		}

		// Create the directory structure for the package.
		err := os.MkdirAll(classPath+"/"+packageName, 0755)
		if err != nil {
			panic(err)
		}

		// Set the current API
		view.API = api
		view.Package = packageName

		// Write the file using our template.
		WriteTemplate(
			packagePath+"/cmd/generate/templates/class.go.tmpl",
			packagePath+"/godot/class/"+packageName+"/"+outFileName,
			view,
		)

		// Run gofmt on the generated Go file.
		log.Println("  Running gofmt on output:", outFileName+"...")
		GoFmt(packagePath + "/godot/class/" + packageName + "/" + outFileName)

		log.Println("  Running goimports on output:", outFileName+"...")
		GoImports(packagePath + "/godot/class/" + packageName + "/" + outFileName)

	}
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
