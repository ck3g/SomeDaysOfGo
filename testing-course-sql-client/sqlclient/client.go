package sqlclient

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type client struct {
	db *sql.DB
}

type row struct {
}

// SQLClient provides an interface of SQL Client
type SQLClient interface {
	Query(query string, args ...interface{}) (rows, error)
}

// Open creates a connection with provided driver name
func Open(driverName, dataSourceString string) (SQLClient, error) {
	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}

	database, err := sql.Open(driverName, dataSourceString)
	if err != nil {
		return nil, err
	}

	c := client{
		db: database,
	}

	return c, nil
}

func (c client) Query(query string, args ...interface{}) (rows, error) {
	returnedRows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	result := sqlRows{
		rows: returnedRows,
	}

	return &result, nil
}
