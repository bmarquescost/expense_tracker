package main

import (
	"github.com/bmarquescost/expense-tracker/db"
	"github.com/bmarquescost/expense-tracker/logger"
	"github.com/bmarquescost/expense-tracker/routes"
	"github.com/bmarquescost/expense-tracker/config"
	"github.com/bmarquescost/expense-tracker/migrations"
	"github.com/joho/godotenv"
	"github.com/gorilla/handlers"
	"go.uber.org/zap"
	"net/http"

)

func main() {
	logger.InitLogger()
	defer logger.CloseLogger()

	// Loading environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal("Error loading .env file", zap.Error(err))
	}

	db, err := db.ConnectDB()
	if err != nil {
		logger.Logger.Fatal("Error while trying to connect to MySQL database", zap.Error(err))
		panic(err)
	}
	defer db.Close()

	migrations.Migrate(db)

	router := routes.SetupRouter(db) 
	
	cfg := config.LoadConfig()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    originsOk := handlers.AllowedOrigins([]string{"http://localhost:5174"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	logger.Logger.Infof("Starting server on %s", cfg.ServerAddr)
	if err := http.ListenAndServe(cfg.ServerAddr, handlers.CORS(headersOk, originsOk, methodsOk)(router)); err != nil {
		logger.Logger.Fatal("Could not  start server: %s", zap.Error(err))
	}
}