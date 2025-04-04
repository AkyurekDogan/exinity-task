/*
The dto package keeps the data transfer objects as http response or inputs in the http
These structs cab be serialized to JSON so can be used as data transfer objects
*/
package model

// Config represents the YAML structure
type Config struct {
	Server struct {
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
	} `yaml:"database"`
}
