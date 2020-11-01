package architecture

import (
	"fmt"

	models "github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/models"
)

// Accessor is how to strore or retrieve a person
// When retrieving a person, if they do not exist, return the zero value
type Accessor interface {
	Save(n int, p models.Person)
	Retrieve(n int) models.Person
}

type PersonService struct {
	a Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}

func (ps PersonService) Get(n int) (models.Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return models.Person{}, fmt.Errorf("No person with ID of %d", n)
	}

	return p, nil
}

func Put(a Accessor, n int, p models.Person) {
	a.Save(n, p)
}
