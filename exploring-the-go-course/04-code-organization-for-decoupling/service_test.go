package architecture

import (
	"fmt"
	"testing"

	"github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/models"
	"github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/04-code-organization-for-decoupling/storage/mongo"
	gomock "github.com/golang/mock/gomock"
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
func TestPutWithMocks(t *testing.T) {
	ctrl := gomock.NewController(t)
	acc := NewMockAccessor(ctrl)

	p := models.Person{
		First: "John",
	}

	acc.EXPECT().Save(1, p).MinTimes(1)

	Put(acc, 1, p)

	ctrl.Finish()
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
