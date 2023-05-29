package main

import (
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/luckystrike561/vizimind/core/internal/datastore/postgres"
	"github.com/luckystrike561/vizimind/core/internal/server"
	"github.com/luckystrike561/vizimind/core/pkg/regiondo"
	"github.com/rs/zerolog/log"
)

var k = koanf.New(".")

func run() int {
	cfg := &server.Config{}

	// Load JSON config.
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't parse config")

		return 1
	}

	if err := k.Unmarshal("server", cfg); err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't unmarshal config")

		return 1
	}

	// Postgres
	postgresSvc := postgres.New(&postgres.Config{
		Host:     k.String("postgres.host"),
		Port:     k.Int("postgres.port"),
		User:     k.String("postgres.user"),
		Password: k.String("postgres.password"),
		Database: k.String("postgres.database"),
	})
	if err := postgresSvc.Init(); err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't initialize postgres client")

		return 1
	}

	// Regiondo
	regiondoSvc := regiondo.New(&regiondo.Config{
		Debug:      k.Bool("regiondo.debug"),
		URL:        k.String("regiondo.url"),
		Timeout:    k.Duration("regiondo.timeout"),
		PublicKey:  k.String("regiondo.publicKey"),
		PrivateKey: k.String("regiondo.privateKey"),
	})
	if err := regiondoSvc.Init(); err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't initialize regiondo client")

		return 1
	}

	srv := server.New(cfg, postgresSvc, regiondoSvc)
	if err := srv.Init(); err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't initialize server")

		return 1
	}

	if err := srv.Run(); err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't run server")

		return 1
	}

	return 0
}

func main() {
	os.Exit(run())
}
