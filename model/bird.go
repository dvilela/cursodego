package model

// Chicken actions
type Chicken interface {
	Cluck() string
}

// Duck actions
type Duck interface {
	Quack() string
}

// Bird is any bird
type Bird struct {
	Name string `json:"name"`
}

// Cluck is the Portuguese (BR) chicken Onomatopoeia
func (b Bird) Cluck() string {
	return "Cocoricóóó..."
}

// Quack is the Portuguese and English duck Onomatopoeia
func (b Bird) Quack() string {
	return "Quack Quack"
}
