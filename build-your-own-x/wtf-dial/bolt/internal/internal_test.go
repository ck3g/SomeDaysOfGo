package internal_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/ck3g/SomeDaysOfGo/build-your-own-x/wtf-dial/bolt/internal"

	"github.com/ck3g/SomeDaysOfGo/build-your-own-x/wtf-dial"
)

// Ensure dial cna be marshaled and unmarshaled.
func TestMarchalDial(t *testing.T) {
	v := wtf.Dial{
		ID:      "ID",
		Token:   "TOKEN",
		Name:    "MYDIAL",
		Level:   10.2,
		ModTime: time.Now().UTC(),
	}

	var other wtf.Dial
	if buf, err := internal.MarshalDial(&v); err != nil {
		t.Fatal(err)
	} else if err := internal.UnmarshalDial(buf, &other); err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(v, other) {
		t.Fatalf("unexpected copy: %#v", other)
	}
}
