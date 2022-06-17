package model

// Community struct of a community
type Community struct {
	// Name name of a community. E.g. EDteam
	Name string `json:"name"`
}

// Communities slice of communities
type Communities []Community

// Person struct of a person
type Person struct {
	// Name name of a person E. g. Alexys
	Name string `json:"name"`
	// Age age of a person E. g. 40
	Age uint8 `json:"age"`
	// Communities communities a person belogns to
	Communities Communities `json:"communities"`
}

// Persons slice of person
type Persons []Person
