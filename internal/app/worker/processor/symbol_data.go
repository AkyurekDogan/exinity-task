/*
Processor package manages the data processing process with external services.
*/
package processor

import "github.com/AkyurekDogan/exinity-task/internal/app/service"

// SymbolData represents the interface for processing symbol data
type SymbolData interface {
	Process() error
}

type symbolData struct {
	srvSymbolData service.SymbolData
}

// NewSymbolData returns the new symbol data processor
func NewSymbolData(srvSymbolData service.SymbolData) SymbolData {
	return &symbolData{
		srvSymbolData: srvSymbolData,
	}
}

func (s *symbolData) Process() error {
	return nil
}
