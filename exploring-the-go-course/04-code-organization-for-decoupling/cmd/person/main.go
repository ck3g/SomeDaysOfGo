package main

import (
	"fmt"

	"github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/models"
	"github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/storage/mongo"
	"github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/storage/postgres"
)

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

func main() {
	dbm := mongo.DB{}
	dbpg := postgres.DB{}

	p1 := models.Person{"Bob"}
	p2 := models.Person{"John"}

	ps := NewPersonService(dbm)

	Put(dbm, 1, p1)
	Put(dbm, 2, p2)

	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))

	Put(dbpg, 1, p1)
	Put(dbpg, 2, p2)

	ps = NewPersonService(dbpg)
	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))
}
