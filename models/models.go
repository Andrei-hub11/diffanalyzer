package models

type Pessoa struct {
	Nome   string
	Idade  int
	Ativo  bool
	Emails []string
}

type Address struct {
	City    string
	Country string
}

type Profile struct {
	Bio     string
	Tags    []string
	Address Address
}

type Person struct {
	ID      int
	Name    string
	Emails  []string
	Profile Profile
}

type FieldDiff struct {
	Path     string
	Expected interface{}
	Actual   interface{}
}

type DataTypes struct {
	IntValue     int
	Int8Value    int8
	Int16Value   int16
	Int32Value   int32
	Int64Value   int64
	UintValue    uint
	Uint8Value   uint8
	Uint16Value  uint16
	Uint32Value  uint32
	Uint64Value  uint64
	Float32Value float32
	Float64Value float64
	BoolValue    bool
	StringValue  string
}

type MapContainer struct {
	StringMap map[string]string
	IntMap    map[string]int
	NestedMap map[string]map[string]int
	PersonMap map[string]Person
}

type Item struct {
	ID     int
	Status string
	Value  int
}

type ItemCollection struct {
	Items []Item
}
