/*
The main package for the worker service.
This service is responsible for processing symbol data.
*/
package main

import (
	"log"
	"os"

	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/drivers"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/repository"
	"github.com/AkyurekDogan/exinity-task/internal/app/service"
	"github.com/AkyurekDogan/exinity-task/internal/app/worker"
	"github.com/AkyurekDogan/exinity-task/internal/app/worker/processor"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

const (
	// ENV environment file path
	ENV = ".env"
	//ENV_CNF_PATH config path
	ENV_CNF_PATH = "WORKER_CONFIG_PATH"
)

// main entry point
func main() {
	// load environment variables
	err := godotenv.Load(ENV)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// use environment variable to get the config path
	appEnvConfigPath := os.Getenv(ENV_CNF_PATH)
	if appEnvConfigPath == "" {
		log.Fatalf("%s environment variable must be set", ENV_CNF_PATH)
	}
	// unmarshall the config file and get all settings in the configuration file.
	yamlFile, err := os.ReadFile(appEnvConfigPath)
	if err != nil {
		log.Fatalf("Error reading configuration YAML file: %v", err)
	}
	var config worker.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML file: %v", err)
	}
	// initialize db connector
	dbROnlyDriver := drivers.NewPostgres(
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
	)
	// initialize the connection.
	dbROnlyDriverInstance, err := dbROnlyDriver.Init()
	if err != nil {
		log.Fatalf("Could not connect to the database: %s\n", err)
	}
	// initialize repository
	repoSymbolData := repository.NewSymbolData(dbROnlyDriverInstance)
	// initialize services
	srvSymbolData := service.NewSymbolData(repoSymbolData)
	// process & listen the symbol data.
	prcSymbolData := processor.NewSymbolData(srvSymbolData)
	prcSymbolData.Process()
}
