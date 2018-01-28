package government

import (
	"math/rand"
	"strconv"

	"github.com/dvilela/cursodego/model"
)

// RegisterPerson registers a person in the civil files
func RegisterPerson(p *model.Person) {
	p.SetID(strconv.Itoa(rand.Int()))
}
