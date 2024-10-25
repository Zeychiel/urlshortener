package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
	config "urlshortener/configuration"
	"urlshortener/internal/getter"
	"urlshortener/internal/shortener"

	"github.com/caarlos0/env/v6"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	// initLogger()

	db := initDB(cfg)
	// add graceful shutdown
	defer func() {
		_ = db.Close()
	}()

	shortenerUseCase := shortener.NewShortenerUseCase(db)
	getterUseCase := getter.NewGetterUseCase(db)

	router := mux.NewRouter()
	router.Handle("/{shortened_url}", getterUseCase).Methods("GET")
	router.Handle("/", shortenerUseCase).Methods("POST")
	router.Handle("/metrics", promhttp.Handler()).Methods("GET") // TODO: provide a dedicated server in a more compelx project.

	slog.Info("Server is running on port 8000")
	server := &http.Server{
		Addr:              ":8000",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}

	log.Fatal(server.ListenAndServe()) // nolint:gocritic // Well, this should be handled, but for the sake of simplicity,
	// we'll leave it as is
}

func initDB(cfg config.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PGHost, cfg.PGPort, cfg.PGUser, cfg.PGPassword, cfg.PGDB)

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
