/*
Processor package manages the data processing process with external services.
*/
package processor

import (
	"context"
	"encoding/json"

	"github.com/AkyurekDogan/exinity-task/internal/app/model"
	"github.com/AkyurekDogan/exinity-task/internal/app/service"
	"github.com/AkyurekDogan/exinity-task/internal/app/worker/aggregator"
	"go.uber.org/zap"

	"github.com/gorilla/websocket"
)

// SymbolData represents the interface for processing symbol data
type Symbol interface {
	Process(ctx context.Context, conn *websocket.Conn)
}

type symbol struct {
	logger        *zap.SugaredLogger
	srvSymbolData service.SymbolData
	aggregator    aggregator.Aggregator
}

// NewSymbolData returns the new symbol data processor
func NewSymbolData(
	logger *zap.SugaredLogger,
	srvSymbolData service.SymbolData,
	aggregator aggregator.Aggregator,
) Symbol {
	return &symbol{
		logger:        logger,
		srvSymbolData: srvSymbolData,
		aggregator:    aggregator,
	}
}

func (s *symbol) Process(ctx context.Context, conn *websocket.Conn) {
	for {
		select {
		case <-ctx.Done():
			s.logger.Info("Context done, stopping processing")
			return
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				s.logger.Error("Error reading message:", err)
				return
			}
			var sm model.Message
			if err := json.Unmarshal(message, &sm); err != nil {
				s.logger.Error("Unmarshal error:", err)
				continue
			}
			_, oldCandle, err := s.aggregator.Process(ctx, sm.Data)
			if err != nil {
				s.logger.Error("Error processing symbols:", err)
				continue
			}
			if oldCandle != nil {
				// Insert the old candle into the database
				err := s.srvSymbolData.Insert(ctx, *oldCandle)
				if err != nil {
					s.logger.Error("Error inserting old candle:", err)
					continue
				}
				s.logger.Infof("Candle saved: %s", oldCandle.Symbol)
			}
		}
	}
}
