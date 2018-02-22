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
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
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
	ClassMap     map[string]GDAPI
	GodotCalls   map[GodotCallSigKey]GodotCallSignature
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

func (v View) GoValue(returnString string) string {
	return GetGoValue(returnString)
}

// GoValue will convert the Godot value into a valid Go value.
func GetGoValue(returnString string) string {
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

func (v View) IsObjectReturnType(method GDMethod) bool {
	return isClassType(method.ReturnType) && !isEnum(method.ReturnType)
}

func (v View) GodotCall(method GDMethod) string {
	var args []string
	for _, arg := range method.Arguments {
		args = append(args, getICallTypeName(arg.Type))
	}
	icallRetType := getICallTypeName(method.ReturnType)
	sigKey := GodotCallSigKey{icallRetType, strings.Join(args, ",")}

	argNames := []string{"o", strconv.Quote(method.Name)}
	for _, arg := range method.Arguments {
		argName := v.GoArgName(arg.Name)
		if isClassType(arg.Type) && !(arg.Type == "Object") {
			argName = fmt.Sprintf("&%s.Object", argName)
		}
		argNames = append(argNames, argName)
		//TODO: escape_cpp(arg["name"]) + (".ptr()" if isReferenceType(arg["type"]) else "")
	}

	var sig GodotCallSignature
	sig, found := v.GodotCalls[sigKey]
	if !found {
		godotArgs := make([]GodotCallArg, len(args))
		for idx, arg := range args {
			godotArgs[idx] = GodotCallArg{
				fmt.Sprintf("arg%d", idx),
				arg,
			}
		}
		sig = GodotCallSignature{icallRetType, godotArgs}
		v.GodotCalls[sigKey] = sig
	}
	icallName := sig.GodotCallName()
	return fmt.Sprintf("%s(%s)", icallName, strings.Join(argNames, ","))
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
	pkgPath := os.Getenv("PKG_PATH")
	if pkgPath == "" {
		panic("$PKG_PATH environment variable was not set. Be sure to run this from generate.sh!")
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
	// TODO: combine SingletonMap and ClassMap probably
	view.ClassMap = map[string]GDAPI{}
	for _, api := range view.APIs {
		view.SingletonMap[api.Name] = api.Singleton
		view.ClassMap[api.Name] = api
	}

	view.GodotCalls = make(map[GodotCallSigKey]GodotCallSignature)
	// Sort the APIs so they will be generated in order.
	sort.Sort(ByName(view.APIs))
	for _, api := range view.APIs {
		sort.Sort(ByMethodName(api.Methods))
		sort.Sort(ByEnumName(api.Enums))
		sort.Sort(ByPropertyName(api.Properties))
		sort.Sort(BySignalName(api.Signals))
	}

	// Create a template from our template files.
	t, err := template.ParseGlob(filepath.Join(templatePath, "*.template"))
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}
	templateBuffer := bytes.NewBufferString("")

	err = t.ExecuteTemplate(templateBuffer, "classes.go.template", view)
	if err != nil {
		log.Fatal("Unable to apply classes template:", err)
	}

	classesOutputFile := filepath.Join(pkgPath, "godot", "classes.go")
	// Output our template to file
	if err := ioutil.WriteFile(classesOutputFile, templateBuffer.Bytes(), 0644); err != nil {
		log.Fatalf("Unable to write %s: %s", classesOutputFile, err)
	}

	templateBuffer.Reset()
	err = t.ExecuteTemplate(templateBuffer, "godotcalls.go.template", view)
	if err != nil {
		log.Fatal("Unable to apply godotcalls template:", err)
	}

	godotCallsOutputFile := filepath.Join(pkgPath, "godot", "godotcalls.go")
	// Output our template to file
	if err := ioutil.WriteFile(godotCallsOutputFile, templateBuffer.Bytes(), 0644); err != nil {
		log.Fatalf("Unable to write %s: %s", godotCallsOutputFile, err)
	}

	// iCalls := make(map[GodotCallSigKey]struct{})
	// usedClasses := getUsedClasses(view.APIs[2])

	// fmt.Fprintln(os.Stderr, generateClassImplementation(iCalls, usedClasses, view.ClassMap, view.APIs[2]))
	// fmt.Fprintln(os.Stderr, generateICallImplementation(iCalls))

}

type GodotCallSigKey struct {
	ReturnType          string
	ArgTypesStringified string
}

type GodotCallSignature struct {
	ReturnType string
	Arguments  []GodotCallArg
}

type GodotCallArg struct {
	Name string
	Type string
}

func (gcs GodotCallSignature) GetReturnType() string {
	return returnType(gcs.ReturnType)
}

func (gcs GodotCallSignature) GodotCallName() string {
	nameParts := []string{fmt.Sprintf("godotCall_%s", stripName(gcs.ReturnType))}
	for _, arg := range gcs.Arguments {
		nameParts = append(nameParts, stripName(arg.Type))
	}
	return casee.ToCamelCase(strings.Join(nameParts, "_"))
}

func (gcs GodotCallSignature) GodotCallDef() string {
	params := []string{"o Class, methodName string"}
	for _, arg := range gcs.Arguments {
		params = append(params, fmt.Sprintf("%s %s", arg.Name, GetGoValue(arg.Type)))
	}
	return fmt.Sprintf("%s(%s) %s", gcs.GodotCallName(), strings.Join(params, ","), GetGoValue(gcs.ReturnType))
}

var setMember struct{}

func generateClassImplementation(
	icalls map[GodotCallSigKey]struct{},
	usedClasses map[string]struct{},
	classes map[string]GDAPI,
	c GDAPI) string {

	//className := stripName(c.Name)

	var source []string

	// TODO: per -type common stuff
	// source = append(source,"#include <" + class_name + ".hpp>")
	// source = append(source,"")
	// source = append(source,"")

	// source = append(source,"#include <core/GodotGlobal.hpp>")
	// source = append(source,"#include <core/CoreTypes.hpp>")
	// source = append(source,"#include <core/Ref.hpp>")

	// source = append(source,"#include <core/Godot.hpp>")
	// source = append(source,"")

	// source = append(source,"#include \"__icalls.hpp\"")
	// source = append(source,"")
	// source = append(source,"")

	for usedClass, _ := range usedClasses {
		fmt.Fprintf(os.Stderr, "UsedClass: '%s'\n", usedClass)
		if !isEnum(usedClass) {
			source = append(source, "// import usedClass somehow") /// TODO: ref used class "#include <" + strip_name(used_class) + ".hpp>")
		}
	}

	source = append(source, "", "")

	//source = append(source,"namespace godot {")

	//core_object_name = ("___static_object_" + strip_name(c["name"])) if c["singleton"] else "this"

	// source = append(source,"")
	// source = append(source,"")

	// if c["singleton"]:
	//     source = append(source,"static godot_object *" + core_object_name + ";")
	//     source = append(source,"")
	//     source = append(source,"")

	//     # FIXME Test if inlining has a huge impact on binary size
	//     source = append(source,"static inline void ___singleton_init()")
	//     source = append(source,"{")
	//     source = append(source,"\tif (" + core_object_name + " == nullptr) {")
	//     source = append(source,"\t\t" + core_object_name + " = godot::api->godot_global_get_singleton((char *) \"" + strip_name(c["name"]) + "\");")
	//     source = append(source,"\t}")
	//     source = append(source,"}")

	//     source = append(source,"")
	//     source = append(source,"")

	// if c["instanciable"]:
	//     source = append(source,"void *" + strip_name(c["name"]) + "::operator new(size_t)")
	//     source = append(source,"{")
	//     source = append(source,"\treturn godot::api->godot_get_class_constructor((char *)\"" + c["name"] + "\")();")
	//     source = append(source,"}")

	//     source = append(source,"void " + strip_name(c["name"]) + "::operator delete(void *ptr)")
	//     source = append(source,"{")
	//     source = append(source,"\tgodot::api->godot_object_destroy((godot_object *)ptr);")
	//     source = append(source,"}")

	for _, method := range c.Methods {
		methodSig := ""

		fmt.Fprintf(os.Stderr, "Process method '%s'\n", method.Name)
		methodSig += makeGDNativeType(method.ReturnType, classes)
		// TODO: method_signature += stripName(c.Name) + "::" + escape_cpp(method["name"]) + "("

		var argDecls []string
		for _, argument := range method.Arguments {
			fmt.Fprintf(os.Stderr, "Process argument '%s':%s\n", argument.Name, argument.Type)
			argDecls = append(argDecls, argument.Name+" "+makeGDNativeType(argument.Type, classes))
		}
		methodSig += strings.Join(argDecls, ", ")

		if method.HasVarargs {
			if len(method.Arguments) > 0 {
				methodSig += ", "
			}
			methodSig += "const Array& __var_args" // TODO:
		}

		// TODO: methodSig += ")" + (" const" if method["is_const"] and not c["singleton"] else "")

		source = append(source, methodSig+" {")

		// if c["singleton"]:
		//     source = append(source,"\t___singleton_init();")

		source = append(source, fmt.Sprintf("mb *C.godot_method_bind = godot_method_bind_get_method(%s, %s);", c.Name, method.Name))

		returnStatement := ""

		if method.ReturnType != "void" {
			// TODO: handle return type
			if isClassType(method.ReturnType) {
				if isEnum(method.ReturnType) {
					returnStatement += "return (" + removeEnumPrefix(method.ReturnType) + ") "
				} else if isReferenceType(method.ReturnType, classes) {
					// TODO: handle references
					returnStatement += "return Ref<" + stripName(method.ReturnType) + ">::__internal_constructor("
				} else {
					ret := ""
					if isClassType(method.ReturnType) {
						ret = "(" + stripName(method.ReturnType) + " *) "
					}
					returnStatement += "return " + ret
				}
			} else {
				returnStatement += "return "
			}
		}

		if method.HasVarargs {

			//TODO: handle vararg variant calls

			// if len(method["arguments"]) != 0:
			//     source = append(source,"\tVariant __given_args[" + str(len(method["arguments"])) + "];")

			// for i, argument in enumerate(method["arguments"]):
			//     source = append(source,"\tgodot::api->godot_variant_new_nil((godot_variant *) &__given_args[" + str(i) + "]);")

			// source = append(source,"")

			// for i, argument in enumerate(method["arguments"]):
			//     source = append(source,"\t__given_args[" + str(i) + "] = " + escape_cpp(argument["name"]) + ";")

			// source = append(source,"")

			// size = ""
			// if method["has_varargs"]:
			//     size = "(__var_args.size() + " + str(len(method["arguments"])) + ")"
			// else:
			//     size = "(" + str(len(method["arguments"])) + ")"

			// source = append(source,"\tgodot_variant **__args = (godot_variant **) alloca(sizeof(godot_variant *) * " + size + ");")

			// source = append(source,"")

			// for i, argument in enumerate(method["arguments"]):
			//     source = append(source,"\t__args[" + str(i) + "] = (godot_variant *) &__given_args[" + str(i) + "];")

			// source = append(source,"")

			// if method["has_varargs"]:
			//     source = append(source,"\tfor (int i = 0; i < __var_args.size(); i++) {")
			//     source = append(source,"\t\t__args[i + " + str(len(method["arguments"])) + "] = (godot_variant *) &((Array &) __var_args)[i];")
			//     source = append(source,"\t}")

			// source = append(source,"")

			// source = append(source,"\tVariant __result;")
			// source = append(source,"\t*(godot_variant *) &__result = godot::api->godot_method_bind_call(mb, (godot_object *) " + core_object_name + ", (const godot_variant **) __args, " + size + ", nullptr);")

			// source = append(source,"")

			// for i, argument in enumerate(method["arguments"]):
			//     source = append(source,"\tgodot::api->godot_variant_destroy((godot_variant *) &__given_args[" + str(i) + "]);")

			// source = append(source,"")

			// if method["return_type"] != "void":
			//     cast = ""
			//     if is_class_type(method["return_type"]):
			//         if isReferenceType(method["return_type"]):
			//             cast += "Ref<" + strip_name(method["return_type"]) + ">::__internal_constructor(__result);"
			//         else:
			//             cast += "(" + strip_name(method["return_type"]) + " *) (Object *) __result;"
			//     else:
			//         cast += "__result;"
			//     source = append(source,"\treturn " + cast)

		} else {

			var args []string
			for _, arg := range method.Arguments {
				args = append(args, getICallTypeName(arg.Type))
			}

			icallRetType := getICallTypeName(method.ReturnType)
			sig := GodotCallSigKey{icallRetType, strings.Join(args, ",")}
			icalls[sig] = setMember
			icallName := "whatevs" // TODO:getICallNamegetICallName(sig)

			argNames := []string{"o"}
			for _, arg := range method.Arguments {
				// TODO: escape Go reserved words
				argNames = append(argNames, arg.Name)
				//TODO: escape_cpp(arg["name"]) + (".ptr()" if isReferenceType(arg["type"]) else "")
			}
			returnStatement += fmt.Sprintf("%s(%s)", icallName, strings.Join(argNames, ","))

			source = append(source, "\t"+returnStatement) // TODO (")" if isReferenceType(method.ReturnType else "") + ";")
		}

		source = append(source, "", "")
	}

	source = append(source, "}")
	return strings.Join(source, "\n")
}

func returnType(t string) string {
	if isClassType(t) {
		return "C.godot_object"
	}

	if t == "int" {
		return "C.godot_int"
	}

	if t == "float" || t == "real" {
		return "C.godot_real"
	}

	if t == "bool" {
		return "C.godot_bool"
	}

	coreType, found := coreTypesGo2Godot[t]
	if found {
		return coreType
	}

	return t
}

func getICallTypeName(name string) string {
	if isEnum(name) {
		return "int"
	}
	if isClassType(name) {
		return "Object" //TODO: right name
	}
	return name
}

func isReferenceType(t string, classes map[string]GDAPI) bool {
	c, found := classes[t]
	if !found {
		return false
	}
	return c.IsReference
}

func makeGDNativeType(t string, classes map[string]GDAPI) string {
	fmt.Fprintf(os.Stderr, "Making GDNativeType for '%s'\n", t)
	if isEnum(t) {
		return removeEnumPrefix(t) + " "
	}

	if isClassType(t) {
		// TODO fix this CPP shit
		if isReferenceType(t, classes) {
			return "Ref<" + stripName(t) + "> "
		}
		return stripName(t) + " *"
	}

	if t == "int" {
		return "int64_t "
	}

	if t == "float" || t == "real" {
		return "double "
	}

	return stripName(t) + " "
}

func isEnum(name string) bool {
	return strings.Index(name, "enum.") == 0
}

func stripName(name string) string {
	if name[0] == '_' {
		return name[1:]
	}
	return name
}

func removeEnumPrefix(name string) string {
	idx := strings.Index(name, "enum.")
	return stripName(name[idx+5:])
}

func getUsedClasses(g GDAPI) map[string]struct{} {
	classes := make(map[string]struct{})
	for _, method := range g.Methods {
		if isClassType(method.ReturnType) {
			if _, found := classes[method.ReturnType]; !found {
				classes[method.ReturnType] = setMember
			}
		}

		for _, arg := range method.Arguments {
			if isClassType(arg.Type) {
				if _, found := classes[arg.Type]; !found {
					classes[arg.Type] = setMember
				}
			}
		}
	}
	return classes
}

func isClassType(name string) bool {
	return !(isCoreType(name) || isPrimitive(name))
}

var coreTypesGo2Godot = map[string]string{
	"Array":            "C.godot_array",
	"Basis":            "C.godot_basis",
	"Color":            "C.godot_color",
	"Dictionary":       "C.godot_dictionary",
	"Error":            "C.godot_error",
	"NodePath":         "C.godot_node_path",
	"Plane":            "C.godot_plane",
	"PoolByteArray":    "C.godot_pool_byte_array",
	"PoolIntArray":     "C.godot_pool_int_array",
	"PoolRealArray":    "C.godot_pool_real_array",
	"PoolStringArray":  "C.godot_pool_string_array",
	"PoolVector2Array": "C.godot_pool_vector2_array",
	"PoolVector3Array": "C.godot_pool_vector3_array",
	"PoolColorArray":   "C.godot_pool_color_array",
	"Quat":             "C.godot_quat",
	"Rect2":            "C.godot_rect2",
	"AABB":             "C.godot_aabb",
	"RID":              "C.godot_rid",
	"String":           "C.godot_string",
	"Transform":        "C.godot_transform",
	"Transform2D":      "C.godot_transform2d",
	"Variant":          "C.godot_variant",
	"Vector2":          "C.godot_vector2",
	"Vector3":          "C.godot_vector3",
}

func isCoreType(name string) bool {
	_, ok := coreTypesGo2Godot[name]
	return ok
}

var primitiveTypes = map[string]struct{}{
	"int":   setMember,
	"bool":  setMember,
	"real":  setMember,
	"float": setMember,
	"void":  setMember,
}

func isPrimitive(name string) bool {
	_, ok := primitiveTypes[name]
	return ok
}

// generate icall functions

func generateICallImplementation(icalls map[GodotCallSigKey]struct{}) string {
	var source []string
	source = append(source, "")

	source = append(source, "#include <gdnative_api_struct.gen.h>")
	source = append(source, "#include <stdint.h>")
	source = append(source, "")

	source = append(source, "")
	source = append(source, "")

	source = append(source, "")

	for icall, _ := range icalls {
		methodSig := ""

		retType := icall.ReturnType
		var args []string
		if icall.ArgTypesStringified != "" {
			args = strings.Split(icall.ArgTypesStringified, ",")
		}

		var methodArgs []string
		for i, arg := range args {
			// TODO: handle this distinction in Go

			// if isCoreType(arg) {
			// 	methodSig += arg + "& "
			// } else
			methodArg := arg
			if !isPrimitive(arg) {
				methodArg = "*Object"
			}

			methodArgs = append(methodArgs, fmt.Sprintf("arg%d %s", i, methodArg))
		}

		//methodSig += fmt.Sprintf("func %s(%s) %s", getICallName(icall), strings.Join(methodArgs, ","), returnType(retType))

		source = append(source, methodSig+" {")

		if retType != "void" {
			source = append(source, "\t"+returnType(retType)+"ret;")
			if isClassType(retType) {
				source = append(source, "\tret = nullptr;")
			}
		}

		arrIdx := ""
		if len(args) == 0 {
			arrIdx = "1"
		}
		source = append(source, "\tconst void *args["+arrIdx+"] = {")

		for i, arg := range args {

			wrappedArg := "\t\t"
			if isPrimitive(arg) || isCoreType(arg) {
				wrappedArg += "(void *) &arg" + string(i)
			} else {
				wrappedArg += "(void *) arg" + string(i)
			}

			wrappedArg += ","
			source = append(source, wrappedArg)
		}

		source = append(source, "\t};")
		source = append(source, "")

		// handle return type
		callRetParam := "&ret"
		if retType == "void" {
			callRetParam = "nullptr"
		}
		source = append(source, "\tgodot::api->godot_method_bind_ptrcall(mb, inst, args, "+callRetParam+");")

		if retType != "void" {
			source = append(source, "\treturn ret;")
		}

		source = append(source, "}")
	}

	source = append(source, "}")
	source = append(source, "")

	return strings.Join(source, "\n")
}
