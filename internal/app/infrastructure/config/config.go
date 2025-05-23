/*
Configuration file for Api Config
This file contains the structure of the configuration file
It is used to read the configuration file and unmarshal it into a struct
It is used to set up the database connection and other settings
It is used to set up the server and other settings
It is used to set up the logger and other settings
It is used to set up the environment variables and other settings
*/
package config

// Worker represents the YAML structure
type Config struct {
	Provider struct {
		Binance struct {
			Scheme    string   `yaml:"scheme"`
			Host      string   `yaml:"host"`
			Path      string   `yaml:"path"`
			Key       string   `yaml:"key"`
			Separator string   `yaml:"separator"`
			Symbols   []string `yaml:"symbols"`
		} `yaml:"binance"`
	} `yaml:"provider"`
	Service struct {
		Network string `yaml:"network"`
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
	} `yaml:"service"`
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
	} `yaml:"database"`
}
