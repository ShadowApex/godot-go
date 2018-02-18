#include <gdnative/string.h>
#include <gdnative_api_struct.gen.h>

#ifndef CGDNATIVE_H
#define CGDNATIVE_H
void godot_go_print_warning(godot_gdnative_core_api_struct *p_api,
			    const char *p_description, const char *p_function,
			    const char *p_file, int p_line);

void godot_go_print_error(godot_gdnative_core_api_struct *p_api,
			  const char *p_description, const char *p_function,
			  const char *p_file, int p_line);

void godot_go_string_destroy(godot_gdnative_core_api_struct *p_api,
			     godot_string *p_self);

void godot_go_print(godot_gdnative_core_api_struct *p_api,
		    const godot_string *p_message);

void godot_go_string_new_with_wide_string(godot_gdnative_core_api_struct *p_api,
					  godot_string *r_dest,
					  const wchar_t *p_contents,
					  const int p_size);
#endif
