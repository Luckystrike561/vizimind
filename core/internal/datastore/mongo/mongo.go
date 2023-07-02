package mongo

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connection = "mongodb://%s:%s@%s:%d/?authSource=admin"
	database   = "vizimind"
	collection = "core"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
}

type Mongo struct {
	cfg        *Config
	collection *mongo.Collection
}

func New(cfg *Config) *Mongo {
	return &Mongo{
		cfg: cfg,
	}
}

func (m *Mongo) Init() error {
	// Set up MongoDB connection options
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf(connection, m.cfg.User, m.cfg.Password, m.cfg.Host, m.cfg.Port),
	)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Error().Err(err).Msg("couldn't connect to mongo")

		return err
	}

	// Ping the MongoDB server to check if the connection is successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("couldn't connect to mongo")

		return err
	}

	log.Info().Msg("connected to mongo")

	m.collection = client.Database(database).Collection(collection)

	return nil
}
