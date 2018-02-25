package methods

type Method struct {
	Arguments      []Argument
	GoName         string
	GoReturns      string
	Name           string
	ParentStruct   string
	Returns        string
	ReturnsPointer bool
}

type Argument struct {
	GoName    string
	GoType    string
	IsPointer bool
	Name      string
	Type      string
}
