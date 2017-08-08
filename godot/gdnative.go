package godot

// GodotGDNativeInitOptions is a wrapper for the `godot_gdnative_init_options`
// structure defined in gdnative.h.
type GodotGDNativeInitOptions struct {
	InEditor      bool
	CoreAPIHash   uint64
	EditorAPIHash uint64
	NoAPIHash     uint64
}

// GodotGDNativeInit is a function signature for functions that will be called
// upon library initialization.
type GodotGDNativeInit func(options *GodotGDNativeInitOptions)

// godotGDNativeInit is a user defined function that will be set by SetGodotGDNativeInit.
var godotGDNativeInit GodotGDNativeInit

// SetGodotGDNativeInit takes an initialization function that will be called
// when Godot loads the Go library.
func SetGodotGDNativeInit(init GodotGDNativeInit) GodotGDNativeInit {
	godotGDNativeInit = init
	return init
}
