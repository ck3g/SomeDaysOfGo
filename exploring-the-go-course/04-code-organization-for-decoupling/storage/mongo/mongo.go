package mongo

import (
	models "github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/models"
)

type DB map[int]models.Person

func (m DB) Save(n int, p models.Person) {
	m[n] = p
}

func (m DB) Retrieve(n int) models.Person {
	return m[n]
}
