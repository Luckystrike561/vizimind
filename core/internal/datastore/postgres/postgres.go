package postgres

import (
	"fmt"

	"github.com/luckystrike561/vizimind/core/internal/model"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const gpxTable = "gpx"

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Postgres struct {
	cfg *Config
	db  *gorm.DB
}

func New(cfg *Config) *Postgres {
	return &Postgres{
		cfg: cfg,
	}
}

func (p *Postgres) Init() error {
	var err error

	p.db, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
			p.cfg.Host, p.cfg.Port, p.cfg.User, p.cfg.Password, p.cfg.Database,
		),
	), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to postgres")

		return err
	}

	if err := p.db.AutoMigrate(&model.Activity{}); err != nil {
		log.Error().Err(err).Msg("failed to auto migrate")

		return err
	}

	return nil
}
