package main

import (
	"log"
	"net/http"
	"os"

	"github.com/A-Victory/user-mig/user/db"
	"github.com/A-Victory/user-mig/user/routes"
	"github.com/A-Victory/user-mig/user/routes/handler"
	"github.com/A-Victory/user-mig/user/service"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	config := os.Getenv("DB_CONFIG")
	db_name := os.Getenv("DB_NAME")

	dbConn := db.NewDBConn(config, db_name)
	defer dbConn.DB.Close()

	svc := service.NewServiceConn(dbConn)

	_ = handler.NewHandlers(svc)

	httpRouter := routes.Router(svc)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", httpRouter); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
