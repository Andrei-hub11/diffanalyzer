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
