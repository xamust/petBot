package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	client     *mongo.Client
	Collection *mongo.Collection
	config     *Config
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {

	client, err := mongo.NewClient(options.Client().ApplyURI(s.config.ConnString))
	if err != nil {
		return err
	}
	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		return err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	s.client = client
	return nil
}

func (s *Store) Disconnect() error {
	return s.client.Disconnect(context.TODO())
}

func (s *Store) GetCollection() *mongo.Collection {
	return s.client.Database(s.config.DBName).Collection(s.config.CollectionName)
}
