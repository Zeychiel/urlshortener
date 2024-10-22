package main

import (
	"database/sql"
	"fmt"
	"net/http"
	config "urlshortener/configuration"
	"urlshortener/internal/shortener"

	"github.com/caarlos0/env/v11"
	_ "github.com/lib/pq"
)

func main() {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	db := initDb(cfg)

	shortenerUseCase := shortener.NewShortenerUseCase(db)

	http.Handle("/", shortenerUseCase)
	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}

func initDb(cfg config.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PG_HOST, cfg.PG_PORT, cfg.PG_USER, cfg.PG_PASSWORD, cfg.PG_DB)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
