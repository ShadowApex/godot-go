package classes

// GDAPIDoc is a structure for parsed documentation.
type GDAPIDoc struct {
	Name        string        `xml:"name,attr"`
	Description string        `xml:"description"`
	Methods     []GDMethodDoc `xml:"methods>method"`
}

type GDMethodDoc struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description"`
}

// GDAPI is a structure for parsed JSON from godot_api.json.
type GDAPI struct {
	APIType      string           `json:"api_type"`
	BaseClass    string           `json:"base_class"`
	Constants    map[string]int64 `json:"constants"`
	Enums        []GDEnums        `json:"enums"`
	Methods      []GDMethod       `json:"methods"`
	Name         string           `json:"name"`
	Properties   []GDProperty     `json:"properties"`
	Signals      []GDSignal       `json:"signals"`
	Singleton    bool             `json:"singleton"`
	Instanciable bool             `json:"instanciable"`
	IsReference  bool             `json:"is_reference"`
}

// ByName is used for sorting GDAPI objects by name
type ByName []GDAPI

func (c ByName) Len() int           { return len(c) }
func (c ByName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByName) Less(i, j int) bool { return c[i].Name < c[j].Name }

type GDEnums struct {
	Name   string           `json:"name"`
	Values map[string]int64 `json:"values"`
}

// ByEnumName is used for sorting GDAPI objects by name
type ByEnumName []GDEnums

func (c ByEnumName) Len() int           { return len(c) }
func (c ByEnumName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByEnumName) Less(i, j int) bool { return c[i].Name < c[j].Name }

type GDMethod struct {
	Arguments    []GDArgument `json:"arguments"`
	HasVarargs   bool         `json:"has_varargs"`
	IsConst      bool         `json:"is_const"`
	IsEditor     bool         `json:"is_editor"`
	IsFromScript bool         `json:"is_from_script"`
	IsNoscript   bool         `json:"is_noscript"`
	IsReverse    bool         `json:"is_reverse"`
	IsVirtual    bool         `json:"is_virtual"`
	Name         string       `json:"name"`
	ReturnType   string       `json:"return_type"`
}

// ByMethodName is used for sorting GDAPI objects by name
type ByMethodName []GDMethod

func (c ByMethodName) Len() int           { return len(c) }
func (c ByMethodName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByMethodName) Less(i, j int) bool { return c[i].Name < c[j].Name }

type GDArgument struct {
	DefaultValue    string `json:"default_value"`
	HasDefaultValue bool   `json:"has_default_value"`
	Name            string `json:"name"`
	Type            string `json:"type"`
}

type GDProperty struct {
	Getter string `json:"getter"`
	Name   string `json:"name"`
	Setter string `json:"setter"`
	Type   string `json:"type"`
}

// ByPropertyName is used for sorting GDAPI objects by name
type ByPropertyName []GDProperty

func (c ByPropertyName) Len() int           { return len(c) }
func (c ByPropertyName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByPropertyName) Less(i, j int) bool { return c[i].Name < c[j].Name }

type GDSignal struct {
	Arguments []GDArgument `json:"arguments"`
	Name      string       `json:"name"`
}

// BySignalName is used for sorting GDAPI objects by name
type BySignalName []GDSignal

func (c BySignalName) Len() int           { return len(c) }
func (c BySignalName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c BySignalName) Less(i, j int) bool { return c[i].Name < c[j].Name }
