/*
The dto package keeps the data transfer objects as http response or inputs in the http
These structs cab be serialized to JSON so can be used as data transfer objects
*/
package dto

// Partner represents the partner data
type Partner struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Loc    Location `json:"location"`
	Radius Measure  `json:"radius"`
	Rating *Rating  `json:"rating"`
	Skills *Skill   `json:"skills"`
}

// Rating represents the rating data
type Rating struct {
	ValueAVG int `json:"value_avg"`
}

// Skill represents the partner skills data
type Skill []string

// Partners represents the partners data
type Partners []Partner
