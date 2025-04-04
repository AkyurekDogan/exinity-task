/*
The drivers package for different db drivers
*/
package drivers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	driver = "postgres"
)

// Driver for database connection. This driver is specificly designed for postgress sql.
type Driver interface {
	Init() (*sql.DB, error)
}

type postgre struct {
	userName string
	password string
	host     string
	port     string
	database string
}

// NewPostgres returns new postgres db manager
func NewPostgres(userName, password, host, port, database string) Driver {
	return &postgre{
		userName: userName,
		password: password,
		host:     host,
		port:     port,
		database: database,
	}
}

// Init creates a new postress instance
func (p *postgre) Init() (*sql.DB, error) {
	// Connection string for PostgreSQL (update with your database credentials)
	connStr := p.getConnectionString()
	// Open the connection
	db, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}
	// Verify connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// getConnectionString creates the connection string
func (p *postgre) getConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", driver, p.userName, p.password, p.host, p.port, p.database)
}
