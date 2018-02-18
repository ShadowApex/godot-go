#include "gdnative.h"
#include <gdnative/string.h>
#include <gdnative_api_struct.gen.h>

void godot_go_print_warning(godot_gdnative_core_api_struct *p_api,
			    const char *p_description, const char *p_function,
			    const char *p_file, int p_line) {
	p_api->godot_print_warning(p_description, p_function, p_file, p_line);
}

void godot_go_print_error(godot_gdnative_core_api_struct *p_api,
			  const char *p_description, const char *p_function,
			  const char *p_file, int p_line) {
	p_api->godot_print_error(p_description, p_function, p_file, p_line);
}

void godot_go_string_destroy(godot_gdnative_core_api_struct *p_api,
			     godot_string *p_self) {
	p_api->godot_string_destroy(p_self);
}

void godot_go_print(godot_gdnative_core_api_struct *p_api,
		    const godot_string *p_message) {
	p_api->godot_print(p_message);
}

void godot_go_string_new_with_wide_string(godot_gdnative_core_api_struct *p_api,
					  godot_string *r_dest,
					  const wchar_t *p_contents,
					  const int p_size) {
	p_api->godot_string_new_with_wide_string(r_dest, p_contents, p_size);
}
