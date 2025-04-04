/*
Configuration file for Api Config
This file contains the structure of the configuration file
It is used to read the configuration file and unmarshal it into a struct
It is used to set up the database connection and other settings
It is used to set up the server and other settings
It is used to set up the logger and other settings
It is used to set up the environment variables and other settings
*/
package api

// Api represents the YAML structure
type Api struct {
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
