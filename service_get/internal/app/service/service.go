package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/xamust/petbot/service_get/internal/app/model"
	"github.com/xamust/petbot/service_get/internal/app/parsefh"
	"github.com/xamust/petbot/service_get/internal/app/store"
	"go.mongodb.org/mongo-driver/bson"
)

type CollectService struct {
	config  *Config
	logger  *logrus.Logger
	parseFH *parsefh.ParserFH
	store   *store.Store
}

// init new server
func New(config *Config) *CollectService {
	return &CollectService{
		config: config,
		logger: logrus.New(),
	}
}

// configure logrus...
func (s *CollectService) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

// config parseFH service...
func (s *CollectService) configureParseFH() error {
	s.parseFH = parsefh.CollectorInit()
	return nil
}

// config store MongoDB....
func (s *CollectService) configStore() error {
	s.store = store.New(s.config.Store)
	if err := s.store.Open(); err != nil {
		return err
	}
	s.store.Collection = s.store.GetCollection()
	return nil
}

func (s *CollectService) StartCollect() error {

	//configure logger...
	if err := s.configureLogger(); err != nil {
		return err
	}

	//configure collecting...
	if err := s.configureParseFH(); err != nil {
		s.logger.Error(err)
		return err
	}

	//configure store...
	if err := s.configStore(); err != nil {
		s.logger.Error(err)
		return err
	}

	//store ops...
	s.logger.Info("Data extract...")
	var result model.ClubData
	filter := bson.D{{"name", "на пулковском"}}

	err := s.store.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	s.logger.Info("Found a single document: %+v\n", result)

	if err = s.store.Disconnect(); err != nil {
		s.logger.Error(err)
		return err
	}

	s.logger.Info("get current FH club success")

	return nil
}
