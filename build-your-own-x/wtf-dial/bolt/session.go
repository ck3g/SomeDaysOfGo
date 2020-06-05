package bolt

import (
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
	// dialService DialService
}
