package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

type GDTypes struct {
	Types []GDType
}

type GDType struct {
	Name       string
	Methods    []GDMethod
	Properties []GDProperty
}

type GDMethod struct {
	Name      string
	Arguments []GDArgument
	Returns   []GDProperty
}

type GDProperty struct {
	Name string
	Type string
}

type GDArgument struct {
	Name string
	Type string
}

func main() {
	// List out template file
	templateFile := "classes.go.template"

	// Create a structure that is used inside our template.
	gdtypes := GDTypes{
		Types: []GDType{
			GDType{
				Name: "Node",
				Methods: []GDMethod{
					GDMethod{
						Name: "get_node",
						Arguments: []GDArgument{
							GDArgument{
								Name: "nodePath",
								Type: "Node",
							},
						},
						Returns: []GDProperty{
							GDProperty{
								Name: "Node",
								Type: "Node",
							},
						},
					},
				},
				Properties: []GDProperty{
					GDProperty{
						Name: "Stuff",
						Type: "string",
					},
				},
			},
		},
	}

	// Create a template from our template file.
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}
	templateBuffer := bytes.NewBufferString("")
	err = t.Execute(templateBuffer, gdtypes)
	if err != nil {
		log.Fatal("Unable to apply email template:", err)
	}

	// Output our template to STDOUT
	fmt.Println(templateBuffer.String())

}
