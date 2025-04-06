/*
Worker Aggregator handles the processing the worker candle data.
*/
package aggregator

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AkyurekDogan/exinity-task/internal/app/model"
	"github.com/AkyurekDogan/exinity-task/internal/app/worker/store"
)

const (
	A_MIN = 60000
)

// Aggregator interface provides interface for candle data processing
type Aggregator interface {
	Process(ctx context.Context, data model.Data) (*model.Candle, *model.Candle, error)
}

type aggregator struct {
	candleStore store.CandleStore
}

// NewAggregator creates a new instance of SymbolData service.
func NewAggregator(
	candleStore store.CandleStore,
) Aggregator {
	return &aggregator{
		candleStore: candleStore,
	}
}

// Process aggregates the candle data to manage.
func (s *aggregator) Process(
	ctx context.Context,
	data model.Data,
) (*model.Candle, *model.Candle, error) {
	// calculate the price
	price, err := parseFloat(data.Price)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing price: %w", err)
	}
	// calculate the quantity
	qty, err := parseFloat(data.Quantity)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing quantity: %w", err)
	}

	// calculate the time slot
	// round down to the nearest minute
	timeSlot := data.EventTime / A_MIN * A_MIN
	// round down to the nearest minute
	// This gives the number of full minutes since epoch, then if we multiply by 60000 we get the time in milliseconds.
	// The purpose it to group trades into 1-minute intervals.
	// For example, if the event time is 123456789 milliseconds since epoch,
	// dividing by 60000 gives 2057613.15, which means 2057613 full minutes since epoch.
	// Multiplying by 60000 gives 123456780000 milliseconds since epoch, which is the start of the minute.
	// This line normalizes the trade timestamp to the start of its 1-minute candle.
	// So if a trade happened at 10:12:34.567, that line gives 10:12:00.000.
	// so we can then group all trades that happen between 10:12:00.000 and 10:12:59.999 into the same OHLC candle.

	candle, exists := s.candleStore.Get(data.Symbol)

	// First time this symbol — create new candle
	if !exists {
		newCandle := &model.Candle{
			Symbol:    data.Symbol,
			OpenTime:  timeSlot,
			Open:      price,
			High:      price,
			Low:       price,
			Close:     price,
			Volume:    qty,
			CloseTime: timeSlot + A_MIN,
		}
		s.candleStore.Update(data.Symbol, newCandle)
		return newCandle, nil, nil // No need to insert yet
	}
	// Existing candle found — check if it's for the current timeSlot
	if candle.OpenTime != timeSlot {
		// Time has moved to next slot — finalize and store the old candle
		// Start a new candle for the new time slot
		newCandle := &model.Candle{
			Symbol:    data.Symbol,
			OpenTime:  timeSlot,
			Open:      price,
			High:      price,
			Low:       price,
			Close:     price,
			Volume:    qty,
			CloseTime: timeSlot + A_MIN,
		}
		s.candleStore.Update(data.Symbol, newCandle)
		// Finalize the old candle
		return newCandle, candle, nil // Return the old candle for insertion
	}
	// Same time slot — update current in-progress candle
	candle.Close = price
	candle.Volume += qty

	if price > candle.High {
		candle.High = price
	}

	if price < candle.Low {
		candle.Low = price
	}
	return nil, nil, nil // Not new not old just update the existing one.
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
