package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/xamust/petbot/service_collect/internal/app/model"
	"github.com/xamust/petbot/service_collect/internal/app/parsefh"
	"github.com/xamust/petbot/service_collect/internal/app/store"
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
	//some ping...
	if _, err := s.parseFH.GetData(); err != nil {
		return err
	}
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
	collectFH, err := s.CollectingToStore()
	if err != nil {
		s.logger.Error(err)
		return err
	}

	insertResult, err := s.store.Collection.InsertMany(context.TODO(), collectFH)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	s.logger.Infof("inserted some data: %v", insertResult.InsertedIDs)

	//for test
	//s.logger.Info("Test extract...")
	//var result model.ClubData
	//filter := bson.D{{"name", "на пулковском"}}
	//
	//err = s.store.Collection.FindOne(context.TODO(), filter).Decode(&result)
	//if err != nil {
	//	s.logger.Error(err)
	//	return err
	//}
	//
	//s.logger.Info("Found a single document: %+v\n", result)

	if err = s.store.Disconnect(); err != nil {
		s.logger.Error(err)
		return err
	}

	s.logger.Info("collect FH clubs success")

	return nil
}

func (s *CollectService) CollectingToStore() ([]interface{}, error) {
	result := make([]interface{}, 0)
	resultMap, err := s.parseFH.GetData()
	if err != nil {
		return nil, err
	}
	for k, v := range resultMap {
		result = append(result, model.ClubData{
			Name: k,
			Url:  v,
		})
	}
	return result, nil
}
