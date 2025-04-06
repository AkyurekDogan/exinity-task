/*
The Entity package for database access and compatible for relevant database
*/
package model

type Message struct {
	Stream string `json:"stream"`
	Data   Data   `json:"data"`
}

type Data struct {
	EType     string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	Price     string `json:"p"`
	Quantity  string `json:"q"`
}
