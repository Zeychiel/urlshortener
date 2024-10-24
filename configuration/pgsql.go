// Package config Configuration for PostgreSQL database.
package config

type Config struct {
	PGDB       string `env:"PG_DB" envDefault:"shortener"`
	PGUser     string `env:"PG_USER" envDefault:"user"`
	PGPassword string `env:"PG_PASSWORD" envDefault:"password"`
	PGHost     string `env:"PG_HOST" envDefault:"localhost"`
	PGPort     int    `env:"PG_PORT" envDefault:"5432"`
}
