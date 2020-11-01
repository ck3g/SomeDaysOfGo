package main

import "fmt"

// Person holds a person structure
type Person struct {
	first string
}

type mongo map[int]Person
type postgres map[int]Person

func (m mongo) save(n int, p Person) {
	m[n] = p
}

func (m mongo) retrieve(n int) Person {
	return m[n]
}

func (pg postgres) save(n int, p Person) {
	pg[n] = p
}

func (pg postgres) retrieve(n int) Person {
	return pg[n]
}

type accessor interface {
	save(n int, p Person)
	retrieve(n int) Person
}

type PersonService struct {
	a accessor
}

func (ps PersonService) get(n int) (Person, error) {
	p := ps.a.retrieve(n)
	if p.first == "" {
		return Person{}, fmt.Errorf("No person with ID of %d", n)
	}

	return p, nil
}

func put(a accessor, n int, p Person) {
	a.save(n, p)
}

func main() {
	dbm := mongo{}
	dbpg := postgres{}

	p1 := Person{"Bob"}
	p2 := Person{"John"}

	ps := PersonService{
		a: dbm,
	}

	put(dbm, 1, p1)
	put(dbm, 2, p2)

	fmt.Println(ps.get(1))
	fmt.Println(ps.get(3))

	put(dbpg, 1, p1)
	put(dbpg, 2, p2)

	ps.a = dbpg
	fmt.Println(ps.get(1))
	fmt.Println(ps.get(3))
}
