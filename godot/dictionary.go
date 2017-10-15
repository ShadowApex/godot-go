package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

//DEPENDS: Array, JSON string
func NewDictionary() *Dictionary {
	dictionary := &Dictionary{}

	// Create our godot dictionary object
	var godotDictionary C.godot_dictionary

	// Create our dictionary
	C.godot_dictionary_new(&godotDictionary)

	// Internally set our dictionary
	dictionary.dictionary = &godotDictionary

	return dictionary
}

// TODO: Finish implementing this
type Dictionary struct {
	dictionary *C.godot_dictionary
}
