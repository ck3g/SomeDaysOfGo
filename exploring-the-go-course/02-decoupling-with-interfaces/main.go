package main

import "fmt"

type person struct {
	first string
}

type mongo map[int]person
type postgres map[int]person

func (m mongo) save(n int, p person) {
	m[n] = p
}

func (m mongo) retrieve(n int) person {
	return m[n]
}

func (pg postgres) save(n int, p person) {
	pg[n] = p
}

func (pg postgres) retrieve(n int) person {
	return pg[n]
}

type accessor interface {
	save(n int, p person)
	retrieve(n int) person
}

func put(a accessor, n int, p person) {
	a.save(n, p)
}

func get(a accessor, n int) person {
	return a.retrieve(n)
}

func main() {
	dbm := mongo{}
	dbpg := postgres{}

	p1 := person{"Bob"}
	p2 := person{"John"}

	put(dbm, 1, p1)
	put(dbm, 2, p2)
	fmt.Println(get(dbm, 1))
	fmt.Println(get(dbm, 2))

	put(dbpg, 1, p1)
	put(dbpg, 2, p2)
	fmt.Println(get(dbpg, 1))
	fmt.Println(get(dbpg, 2))
}
