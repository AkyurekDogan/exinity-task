/*
Repository package for data access
*/
package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AkyurekDogan/exinity-task/internal/app/model"
)

// Symbol represents the interface for symbol data repository
type Symbol interface {
	Insert(ctx context.Context, candle model.Candle) error
}

type symbol struct {
	dbDriver *sql.DB
}

// NewSymbolData creates new database interface for SymbolData
func NewSymbolData(
	dbDriver *sql.DB,
) Symbol {
	return &symbol{
		dbDriver: dbDriver,
	}
}

// Insert inserts the symbolData to database.
func (u *symbol) Insert(ctx context.Context, c model.Candle) error {
	query := `INSERT INTO symbol_data (
		symbol, open_time, open, high, low, close, volume, close_time
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := u.dbDriver.Exec(
		query,
		c.Symbol,
		c.OpenTime,
		c.Open,
		c.High,
		c.Low,
		c.Close,
		c.Volume,
		c.CloseTime,
	)
	if err != nil {
		return fmt.Errorf("error on insert repo: %w", err)
	}
	return nil
}
