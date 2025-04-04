/*
The Entity package for database access and compatible for relevant database
*/
package model

// Partner represents the partner data
type Partner struct {
	Id     string
	Name   string
	Loc    Location
	Radius int
}

// Rating represents the partner ratings data
type Rating struct {
	PartnerId string
	ValueAVG  int
}

// Skill represents the partner skills data
type Skill struct {
	PartnerId string
	Skills    []string
}

// Partners represents the partners data
type Partners []Partner
