/*
The dto package keeps the data transfer objects as http response or inputs in the http
These structs cab be serialized to JSON so can be used as data transfer objects
*/
package dto

// Response represents the http response data
type Response struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

// Success represents the http success response data
type Success struct {
	Response
	Data any `json:"data,omitempty"`
}

// Error represents the http fail response data
type Error struct {
	Response
	Error any `json:"error,omitempty"`
}
