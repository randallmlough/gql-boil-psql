package main

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func getDatabase() *sql.DB {
	// Start database connection
	if os.Getenv("DATABASE_DEBUG") == "true" {
		boil.DebugMode = true
	}

	// Open handle to database like normal
	db, err := sql.Open("postgres", "postgresql://postgres:password@postgres:5432/gqlboil")
	if err != nil {
		log.Fatal().Err(err).Msg("could not open database connection")
	}
	// https://www.alexedwards.net/blog/configuring-sqldb
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("no real database connection")
	}
	return db
}
