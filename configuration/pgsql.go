package config

type Config struct {
	PG_DB       string `env:"PG_DB" envDefault:"shortener"`
	PG_USER     string `env:"PG_USER" envDefault:"user"`
	PG_PASSWORD string `env:"PG_PASSWORD" envDefault:"password"`
	PG_HOST     string `env:"PG_HOST" envDefault:"localhost"`
	PG_PORT     int    `env:"PG_PORT" envDefault:"5432"`
}
