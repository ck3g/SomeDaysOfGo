package sqlclient

import (
	"database/sql"
	"errors"
)

type client struct {
	db *sql.DB
}

type row struct {
}

// SqlClient provides an interface of SQL Client
type SqlClient interface {
	Query(query string, args ...interface{}) (*row, error)
}

// Open creates a connection with provided driver name
func Open(driverName, dataSourceString string) (SqlClient, error) {
	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}

	db, err := sql.Open(driverName, dataSourceString)
	if err != nil {
		return nil, err
	}

	c := client{
		db: db,
	}

	return c, nil
}

func (c client) Query(query string, args ...interface{}) (*row, error) {
	// return c.db.Query(query, args)
	return &row{}, nil
}
