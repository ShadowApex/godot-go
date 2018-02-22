package types

// TypeDef is a structure for holding C type definitions. This is used so we
// can generate a Go wrapper for all Godot base types.
type TypeDef struct {
	Name       string    // The C type name in snake_case
	GoName     string    // The Go type name in camelCase
	HeaderName string    // The header file this type shows up in
	Base       string    // Base will let us know if this is a struct, int, etc.
	Properties []TypeDef // Optional C struct fields
}
