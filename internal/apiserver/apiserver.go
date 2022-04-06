package apiserver

import (
	"net/http"

	"github.com/auth_service/internal/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
	store  *store.Store
}

func New(config *Config, store *store.Store) *ApiServer {
	return &ApiServer{
		config: config,
		store:  store,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
}

func (s *ApiServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Infof("Server starts on port... %s", s.config.Port)

	s.logger.Info("Connection to database...")
	if err := s.store.Open(); err != nil {
		return err
	}

	s.logger.Info("Init handlers:")
	s.initHandlers()

	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *ApiServer) configureLogger() error {
	lvl, err := logrus.ParseLevel(s.config.DebugLvl)
	if err != nil {
		return err
	}

	s.logger.SetLevel(lvl)

	return nil
}
