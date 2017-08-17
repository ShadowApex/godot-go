package godot

// methodMap is a mapping of Go method names to their corresponding method
// in Godot. There may be a better way to do this, but right now it is
// required because public methods cannot start with an underscore or
// lowercase letter.
var methodMap = map[string]string{
	"Ready":   "_ready",
	"Process": "_process",
}
