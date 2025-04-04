/*
The main package for the service
*/
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AkyurekDogan/exinity-task/internal/app/handler"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/drivers"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/middlewares"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/model"
	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/repository"
	"github.com/AkyurekDogan/exinity-task/internal/app/service"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/AkyurekDogan/exinity-task/docs/swagger" // Import Swagger docs

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

const (
	// ENV environment file path
	ENV = ".env"
	//ENV_CNF_PATH config path
	ENV_CNF_PATH = "CONFIG_PATH"
)

// @title Exinity Task
// @version 1.0
// @description This project is build for Exinity take home assessment.
// @contact.name Dogan Akyurek
// @contact.email akyurek.dogan.dgn@gmail.com
// @host localhost:1989
// @BasePath /

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
	var config model.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML file: %v", err)
	}

	// initialize db connector

	// read driver
	dbDriverRead := drivers.NewPostgres(
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
	)

	dbDriverR, err := dbDriverRead.Init()
	if err != nil {
		log.Fatalf("Could not connect to the database: %s\n", err)
	}

	// initialize repository
	dbRepoPartnerSkill := repository.NewSkill(dbDriverR)
	dbRepoPartnerRating := repository.NewRating(dbDriverR)
	dbRepoPartner := repository.NewPartner(dbDriverR)
	dbRepoMatch := repository.NewMatch(dbDriverR)

	// initialize services
	srvPartnerRating := service.NewRating(dbRepoPartnerRating)
	srvPartnerSkill := service.NewSkill(dbRepoPartnerSkill)
	srvPartner := service.NewPartner(dbRepoPartner, srvPartnerSkill, srvPartnerRating)
	matchService := service.NewMatch(dbRepoMatch)

	// handlers
	handlerMatch := handler.NewMatch(matchService)
	handlerPartner := handler.NewPartner(srvPartner)
	// Create a new router
	r := chi.NewRouter()

	r.Use(middlewares.AddHeaderMiddleware())
	// Define the endpoints
	// Swagger UI endpoint
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/match", handlerMatch.Get)
	r.Get("/partner", handlerPartner.Get)

	r.Options("/*", func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for preflight response
		w.Header().Set("Allow", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Path", "/*")
		w.WriteHeader(http.StatusNoContent)
	})

	// Start the HTTP server
	err = http.ListenAndServe(config.Server.Host, r)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
