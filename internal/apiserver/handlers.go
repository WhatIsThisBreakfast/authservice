package apiserver

import "github.com/auth_service/internal/auth"

func (s *ApiServer) initHandlers() {
	authService := auth.New(s.store.Users, s.logger)

	s.logger.Info("Auth handler")
	auth.RegisterRoutes(s.router, authService)

	s.logger.Info("Init complete")
}
