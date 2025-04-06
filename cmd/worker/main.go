/*
The main package for the worker service.
This service is responsible for processing symbol data.
*/
package main

import (
	"context"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/drivers"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/logger"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/repository"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/store"
	"github.com/AkyurekDogan/exinity-task/internal/app/service"
	"github.com/AkyurekDogan/exinity-task/internal/app/worker"
	"github.com/AkyurekDogan/exinity-task/internal/app/worker/aggregator"
	"github.com/AkyurekDogan/exinity-task/internal/app/worker/processor"
	"github.com/gorilla/websocket"

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
	// Initialize structured logger
	log, err := logger.NewLogger()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer log.Sync() // Flush any buffered log entries
	logger := log.Sugar()
	// load environment variables
	err = godotenv.Load(ENV)
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}
	// use environment variable to get the config path
	appEnvConfigPath := os.Getenv(ENV_CNF_PATH)
	if appEnvConfigPath == "" {
		panic("environment variable must be set: " + ENV_CNF_PATH)
	}
	// unmarshall the config file and get all settings in the configuration file.
	yamlFile, err := os.ReadFile(appEnvConfigPath)
	if err != nil {
		panic("error reading configuration YAML file: " + err.Error())
	}
	var config worker.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic("error unmarshalling YAML file: " + err.Error())
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
		panic("Could not connect to the database: " + err.Error())
	}
	// initialize repository
	repoSymbolData := repository.NewSymbolData(dbROnlyDriverInstance)
	// initialize services
	srvSymbolData := service.NewSymbolData(repoSymbolData)
	// in-memory thread-safe storage for symbol data
	symbolCandleStore := store.NewCandleStore()
	// initialize the aggregator
	aggregator := aggregator.NewAggregator(symbolCandleStore)
	// process & listen the symbol data.
	prcSymbolData := processor.NewSymbolData(logger, srvSymbolData, aggregator)
	// prepare the url for the websocket connection
	u := url.URL{
		Scheme: config.Provider.Binance.Scheme,
		Host:   config.Provider.Binance.Host,
		Path:   config.Provider.Binance.Path,
		RawQuery: config.Provider.Binance.Key + "=" +
			strings.Join(config.Provider.Binance.Symbols, config.Provider.Binance.Separator+"/") +
			config.Provider.Binance.Separator,
	}
	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic("WebSocket dial error: " + err.Error())
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	// run in a goroutine
	go prcSymbolData.Process(ctx, conn)
	// Wait for shutdown since it is async process
	logger.Info("the worker is running...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	cancel()
	logger.Info("the worker is shutting down...")
}
