package model

// Person represents a person
type Person struct {

	// Name is the name of the person
	Name string `json:"name"`

	// Age is the age of the person
	Age int `json:"age"`

	// Married represents its civil state
	Married bool `json:"married"`

	// private var
	id string
}

// SetID the id setter
func (p *Person) SetID(id string) {
	p.id = id
}

// GetID the id getter
func (p *Person) GetID() string {
	return p.id
}
