/*
The Entity package for database access and compatible for relevant database
*/
package model

type Candle struct {
	Symbol    string  `json:"symbol"`
	OpenTime  int64   `json:"open_time"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    float64 `json:"volume"`
	CloseTime int64   `json:"close_time"`
}
