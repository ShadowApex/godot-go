package types

// TypeDef is a structure for holding C type definitions. This is used so we
// can generate a Go wrapper for all Godot base types.
type TypeDef struct {
	Base       string    // Base will let us know if this is a struct, int, etc.
	Comment    string    // Contains the comment on the line of the struct
	GoName     string    // The Go type name in camelCase
	HeaderName string    // The header file this type shows up in
	IsPointer  bool      // Usually for properties; defines if it is a pointer type
	Name       string    // The C type name in snake_case
	Properties []TypeDef // Optional C struct fields
	SimpleType bool      // Whether or not the definition is just one line long (e.g. bool, int, etc.)
}
