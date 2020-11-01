package main

import (
	"fmt"

	architecture "github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling"
	"github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/models"
	"github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/storage/mongo"
	"github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/storage/postgres"
)

func main() {
	dbm := mongo.DB{}
	dbpg := postgres.DB{}

	p1 := models.Person{"Bob"}
	p2 := models.Person{"John"}

	ps := architecture.NewPersonService(dbm)

	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)

	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))

	architecture.Put(dbpg, 1, p1)
	architecture.Put(dbpg, 2, p2)

	ps = architecture.NewPersonService(dbpg)
	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))
}
