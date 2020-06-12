package bolt

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/ck3g/SomeDaysOfGo/build-your-own-x/wtf-dial"
)

// Client represents a client to the underlying BoldDB data store.
type Client struct {
	// Filename to the BoltDB database.
	Path string

	// Returns the current time.
	Now func() time.Time

	// Services
	dialService DialService

	db *bolt.DB
}

// NewClient initializes new client.
func NewClient() *Client {
	c := &Client{Now: time.Now}
	c.dialService.client = c
	return c
}

// Open opens and initializes the BoltDB database.
func (c *Client) Open() error {
	// Open database file.
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	c.db = db

	// Initialize top-level buckets.
	tx, err := c.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.CreateBucketIfNotExists([]byte("Dials")); err != nil {
		return err
	}

	return tx.Commit()
}

// Close closes then underlying BoltDB database.
func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

// DialService returns the dial service associated with the client.
func (c *Client) DialService() wtf.DialService { return &c.dialService }
