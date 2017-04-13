# Go-Godot
Go language bindings for the [Godot Engine](https://godotengine.org/)'s [GDNative API](https://github.com/GodotNativeTools/godot_headers).

# Usage
The Godot Engine can interface with code written in Go using the GDNative module. 

# Setup

## Linux
To properly compile your code, you'll need to follow these steps:

* Clone the Godot Headers into a path where the compiler can find them:    
`sudo git clone https://github.com/GodotNativeTools/godot_headers.git /usr/local/include`

# Build
When you're ready to build your code as a dynamic library that can be imported into
Godot, use the following command:    
`go build -buildmode=c-shared -o libgodot.so godot.go`    

## How do I use native scripts from the editor?

*note* the pictures need to be updated to represent the new naming.

Manually create a **GDNativeLibrary** resource.    
![](https://raw.githubusercontent.com/GodotNativeTools/godot_headers/master/images/faq/dllibrary_create_new_resource.png)

![](https://raw.githubusercontent.com/GodotNativeTools/godot_headers/master/images/faq/dllibrary_create_new_dllibrary.png)

Save it as a **.tres**.

![](https://raw.githubusercontent.com/GodotNativeTools/godot_headers/master/images/faq/dllibrary_save_as_resource.png)

This resource contains links to the libraries for each platform.

Now, **create a new native script on your node.** You may prefer built-in script for native scripts, unless you want to organize many **.gdn** files which only contain a name. **You must specify the name of the class you would like to use.**

![](https://raw.githubusercontent.com/GodotNativeTools/godot_headers/master/images/faq/create_dlscript.png?raw=true)

"Native" resources have a field for a GDNativeLibrary. **This field is empty by default.**

![](https://raw.githubusercontent.com/GodotNativeTools/godot_headers/master/images/faq/set_script_dllibrary.png?raw=true)

If you leave it empty, **you can set a default GDNativeLibrary in the project settings** ```dlscript/default_dllibrary```.

![](https://raw.githubusercontent.com/GodotNativeTools/godot_headers/master/images/faq/set_project_dllibrary.png?raw=true)
