/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"errors"

	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/repository"
)

var (
	ErrNoPartner = errors.New("the symbol data is not found")
)

// SymbolData interface provides interface for SymbolData service
type SymbolData interface {
	Get() error
	Insert() error
}

type symbolData struct {
	dbSymbolData repository.SymbolData
}

// NewSymbolData creates a new instance of SymbolData service.
func NewSymbolData(
	repoSymbolData repository.SymbolData,
) SymbolData {
	return &symbolData{
		dbSymbolData: repoSymbolData,
	}
}

// Get returns the relevant symbol data
func (s *symbolData) Get() error {
	return nil
}

// Insert inserts the symbol data into the database
func (s *symbolData) Insert() error {
	return nil
}
