/*
Configuration files for the worker service
This file contains the configuration structure for the worker service.
It includes the database configuration such as username, password, host, port, and database name.
*/
package worker

// Worker represents the YAML structure
type Worker struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
	} `yaml:"database"`
}
