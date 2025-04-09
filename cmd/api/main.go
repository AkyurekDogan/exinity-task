/*
The main package for the worker service.
This service is responsible for processing symbol data.
*/
package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/AkyurekDogan/exinity-task/internal/app/aggregator"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/config"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/drivers"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/logger"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/repository"
	"github.com/AkyurekDogan/exinity-task/internal/app/processor"
	candlepb "github.com/AkyurekDogan/exinity-task/internal/app/proto"
	grpcserver "github.com/AkyurekDogan/exinity-task/internal/app/server"
	"github.com/AkyurekDogan/exinity-task/internal/app/service"
	"github.com/AkyurekDogan/exinity-task/internal/app/store"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

const (
	// ENV environment file path
	ENV = ".env"
	//ENV_CNF_PATH config path
	ENV_CNF_PATH = "CONFIG_PATH"
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
	var config config.Config
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
		panic("Could not connect to the database: " + err.Error() + " Info: " + config.Database.Host + ":" + config.Database.Port)
	}
	// initialize repository
	repoSymbolData := repository.NewSymbolData(dbROnlyDriverInstance)
	// initialize services
	srvSymbolData := service.NewSymbolData(repoSymbolData)
	// in-memory thread-safe storage for symbol data
	symbolCandleStore := store.NewCandleStore()
	// initialize the aggregator
	aggregator := aggregator.NewAggregator(symbolCandleStore)

	// initialize the grpc server
	grpcSrv := grpcserver.NewCandleServiceServer()

	// process & listen the symbol data.
	prcSymbolData := processor.NewSymbolData(logger, srvSymbolData, aggregator, grpcSrv)
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
	// Start gRPC server in a separate goroutine
	go startGRPCServer(config, grpcSrv, logger)
	// Wait for shutdown since it is async process
	logger.Info("the worker is running...")
	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	cancel()
	logger.Info("the worker is shutting down...")
}

func startGRPCServer(
	cnf config.Config,
	grpcSrv *grpcserver.CandleServiceServer,
	logger *zap.SugaredLogger,
) {
	lis, err := net.Listen(cnf.Service.Network, fmt.Sprintf("%s:%s", cnf.Service.Host, cnf.Service.Port)) // gRPC listens on this port
	if err != nil {
		panic("Failed to listen: " + err.Error())
	}
	s := grpc.NewServer()
	candlepb.RegisterCandleServiceServer(s, grpcSrv)

	logger.Infof("gRPC server started on %s:%s", cnf.Service.Host, cnf.Service.Port)
	// Graceful shutdown
	if err := s.Serve(lis); err != nil {
		panic("Failed to serve: " + err.Error())
	}
}
