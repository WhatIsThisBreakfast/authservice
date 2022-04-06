package usersrepository

import (
	"database/sql"

	"github.com/auth_service/internal/models"
)

type UsersRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (r *UsersRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.QueryRow(
		createUserQuery,
		user.PublicID,
		user.Payload,
	).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UsersRepository) GetUser(publicID string) (*models.User, error) {
	user := &models.User{}

	if err := r.db.QueryRow(
		getUserQuery,
		publicID,
	).Scan(&user.ID, &user.PublicID, &user.Payload); err != nil {
		return nil, err
	}

	return user, nil
}
