package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type DBconn struct {
	DB *sql.DB
}

func NewDBConn(config, dbName string) *DBconn {

	db, err := sql.Open("postgres", config)
	if err != nil {
		log.Fatal("unable to open database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("unable to connect to database: ", err)
	}

	if err := goose.Up(db, "user/db/migration"); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return &DBconn{
		DB: db,
	}
}
