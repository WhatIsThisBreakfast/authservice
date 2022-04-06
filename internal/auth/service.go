package auth

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/auth_service/internal/models"
	"github.com/auth_service/internal/repositories/usersrepository"
	"github.com/auth_service/pkg/serviceutil"
	"github.com/sirupsen/logrus"
)

type Service struct {
	users  *usersrepository.UsersRepository
	logger *logrus.Logger
}

func New(users *usersrepository.UsersRepository, logger *logrus.Logger) *Service {
	return &Service{
		users:  users,
		logger: logger,
	}
}

func (s *Service) IsUserExist(publicID string) bool {
	_, err := s.users.GetUser(publicID)

	return err == nil
}

func (s *Service) createNewUser(publicID string, payload string) (*models.User, *serviceutil.ServiceError) {
	hashed_payload := sha256.Sum256([]byte(payload + publicID))
	string_payload := fmt.Sprintf("%x", hashed_payload)

	user := &models.User{
		PublicID: publicID,
		Payload:  string_payload,
	}

	if s.IsUserExist(publicID) {
		return nil, serviceutil.NewError(http.StatusBadRequest, userExists)
	}

	createdUser, err := s.users.CreateUser(user)
	if err != nil {
		s.logger.Error(err)
		return nil, serviceutil.NewError(http.StatusInternalServerError, notCreatedMessage)
	}

	return createdUser, nil
}

func (s *Service) AuthUser(publicID string, payload string) (*models.User, *serviceutil.ServiceError) {
	hashed_payload := sha256.Sum256([]byte(payload + publicID))
	string_payload := fmt.Sprintf("%x", hashed_payload)

	findedUser, err := s.users.GetUser(publicID)
	if err != nil || findedUser.Payload != string_payload {
		return nil, serviceutil.NewError(http.StatusBadRequest, incorrectData)
	}

	return findedUser, nil
}
