package service

import (
	"github.com/sirupsen/logrus"
	ci "github.com/xamust/petbot/service_get/api"
	"github.com/xamust/petbot/service_get/internal/app/parsefh"
	"github.com/xamust/petbot/service_get/internal/app/store"
	"google.golang.org/grpc"
	"net"
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

	//gRPC...
	s.logger.Info("Get data...")

	listengRPS, err := net.Listen("tcp", s.config.PortgRPC)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	serviceCollect := grpc.NewServer()
	ci.RegisterClubsInfoServer(serviceCollect, &server{collectService: s})

	s.logger.Infof("Starting gRPC listener on port: %v ...", s.config.PortgRPC)

	if err = serviceCollect.Serve(listengRPS); err != nil {
		s.logger.Error("failed to serve: %v", err)
		return err
	}

	return nil
}
