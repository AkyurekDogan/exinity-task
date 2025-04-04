/*
The dto package keeps the data transfer objects as http response or inputs in the http
These structs cab be serialized to JSON so can be used as data transfer objects
*/
package dto

// Location struct for coordinates
type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// Measure represents the measure standard data
type Measure struct {
	Value  float32 `json:"value"`
	Metric string  `json:"metric"`
}

// Match struct for matching data
type Match struct {
	PartnerId string   `json:"partner_id"`
	Name      string   `json:"name"`
	Loc       Location `json:"location"`
	Radius    Measure  `json:"radius"`
	Distance  Measure  `json:"distance"`
	Rating    Rating   `json:"rating"`
	Skills    Skill    `json:"skills"`
	Rank      int      `json:"rank"`
}

// MatchList represents the response struct that for success response
type MatchList []Match

// MatchList represents the response struct that for success response
type MatchListResponse struct {
	Filter  MatchFilter `json:"filter"`
	Matches MatchList   `json:"matches"`
}
