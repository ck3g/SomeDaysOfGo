package postgres

import (
	models "github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/models"
)

type DB map[int]models.Person

func (pg DB) Save(n int, p models.Person) {
	pg[n] = p
}

func (pg DB) Retrieve(n int) models.Person {
	return pg[n]
}
