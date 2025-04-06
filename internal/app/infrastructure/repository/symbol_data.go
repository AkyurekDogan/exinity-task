/*
Repository package for data access
*/
package repository

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRows = errors.New("no record found with the given id")
)

// SymbolData represents the interface for symbol data repository
type SymbolData interface {
	Get() error
	Insert() error
}

type symbolData struct {
	dbDriver *sql.DB
}

// NewSymbolData creates new database interface for SymbolData
func NewSymbolData(
	dbDriver *sql.DB,
) SymbolData {
	return &symbolData{
		dbDriver: dbDriver,
	}
}

// Get gets the symbolData from database to models.
func (u *symbolData) Get() error {
	return nil
}

// Insert inserts the symbolData to database.
func (u *symbolData) Insert() error {
	return nil
}
