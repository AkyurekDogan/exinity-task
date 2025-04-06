/*
Service package handles the services for business logic and data processing
*/
package service

import (
	"context"
	"fmt"

	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/repository"
	"github.com/AkyurekDogan/exinity-task/internal/app/model"
)

// SymbolData interface provides interface for SymbolData service
type SymbolData interface {
	Insert(ctx context.Context, c model.Candle) error
}

type symbolData struct {
	dbSymbolData repository.Symbol
}

// NewSymbolData creates a new instance of SymbolData service.
func NewSymbolData(
	repoSymbolData repository.Symbol,
) SymbolData {
	return &symbolData{
		dbSymbolData: repoSymbolData,
	}
}

// Insert inserts the symbol data into the database
func (s *symbolData) Insert(ctx context.Context, c model.Candle) error {
	err := s.dbSymbolData.Insert(ctx, c)
	if err != nil {
		return fmt.Errorf("error in insert service: %w", err)
	}
	return nil
}
