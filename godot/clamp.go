package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

// Clamp will restrict the given value to the range between min and max.
func Clamp(value, min, max gdnative.Real) gdnative.Real {
	if value < min {
		return min
	}
	if value > max {
		return max
	}

	return value
}
