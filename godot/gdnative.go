package godot

// GodotGDNativeInitOptions is a wrapper for the `godot_gdnative_init_options`
// structure defined in gdnative.h.
type GodotGDNativeInitOptions struct {
	InEditor      bool
	CoreAPIHash   uint64
	EditorAPIHash uint64
	NoAPIHash     uint64
}
