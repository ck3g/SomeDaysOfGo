package bolt

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
	"github.com/ck3g/SomeDaysOfGo/build-your-own-x/wtf-dial"
)

// Session represents an authenticable connection to the database.
type Session struct {
	db  *bolt.DB
	now time.Time

	// Authentication
	authenticator wtf.Authenticator
	authToken     string
	user          *wtf.User

	// Services
	dialService DialService
}

// itob returns an 8-byte bit-endian encoded byte slice of v.
//
// This function is typically used for encoding integer IDs to byte slices
// so that they can be used as BoltDB keys.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
