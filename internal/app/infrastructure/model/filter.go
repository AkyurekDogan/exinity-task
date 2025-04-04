/*
The Entity package for database access and compatible for relevant database
*/
package model

// Filter represents the filtering in db compatible mode
type Filter struct {
	PartnerId string
}

// MatchFilter filters the cost details
type MatchFilter struct {
	MaterialType string
	Loc          Location
}
