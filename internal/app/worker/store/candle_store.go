package store

import (
	"sync"

	"github.com/AkyurekDogan/exinity-task/internal/app/model"
)

type CandleStore interface {
	GetOrCreate(symbol string, defaultCandle *model.Candle) *model.Candle
	Update(symbol string, newCandle *model.Candle)
	Get(symbol string) (*model.Candle, bool)
}

type candleStore struct {
	candles sync.Map
}

func NewCandleStore() CandleStore {
	return &candleStore{
		candles: sync.Map{},
	}
}

func (cs *candleStore) GetOrCreate(symbol string, defaultCandle *model.Candle) *model.Candle {
	value, _ := cs.candles.LoadOrStore(symbol, defaultCandle)
	return value.(*model.Candle)
}

func (cs *candleStore) Update(symbol string, newCandle *model.Candle) {
	cs.candles.Store(symbol, newCandle)
}

func (cs *candleStore) Get(symbol string) (*model.Candle, bool) {
	value, ok := cs.candles.Load(symbol)
	if !ok {
		return nil, false
	}
	return value.(*model.Candle), true
}
