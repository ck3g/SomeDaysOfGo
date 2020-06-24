package sqlclient

import (
	"database/sql"
	"errors"
	"os"

	// Needs for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	goEnvironment = "GO_ENVIRONMENT"
	production    = "production"
)

var isMocked bool
var dbClient SQLClient

type client struct {
	db *sql.DB
}

type row struct {
}

// SQLClient provides an interface of SQL Client
type SQLClient interface {
	Query(query string, args ...interface{}) (rows, error)
}

// StartMockServer starts the mock server
func StartMockServer() {
	isMocked = true
}

// StopMockServer stops the mock server
func StopMockServer() {
	isMocked = false
}

func isProduction() bool {
	return os.Getenv(goEnvironment) == production
}

// Open creates a connection with provided driver name
func Open(driverName, dataSourceString string) (SQLClient, error) {
	if isMocked && !isProduction() {
		dbClient = &clientMock{}
		return dbClient, nil
	}

	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}

	database, err := sql.Open(driverName, dataSourceString)
	if err != nil {
		return nil, err
	}

	dbClient = &client{
		db: database,
	}

	return dbClient, nil
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
