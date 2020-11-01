package architecture

import (
	"fmt"
	"testing"

	"github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/models"
	"github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/storage/mongo"
)

func TestPut(t *testing.T) {
	mdb := mongo.DB{}
	p := models.Person{
		First: "John",
	}

	Put(mdb, 1, p)

	got := mdb.Retrieve(1)
	if got != p {
		t.Fatalf("want %v, got %v", p, got)
	}
}

func ExamplePut() {
	mdb := mongo.DB{}
	p := models.Person{
		First: "John",
	}

	Put(mdb, 1, p)
	got := mdb.Retrieve(1)
	fmt.Println(got)
	// Output: {John}
}
