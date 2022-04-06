package store

import (
	"database/sql"

	"github.com/auth_service/internal/repositories/usersrepository"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Store struct {
	config *Config
	db     *sql.DB
	Users  *usersrepository.UsersRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.getDbUrl())
	if err != nil {
		return errors.Wrapf(err, " | \033[31merror on sql.Open. Please, check config\n Connection %s\n", s.config.getDbUrl())
	}

	if err := db.Ping(); err != nil {
		return errors.Wrapf(err, " | \033[31merror on db.Ping. Please, check config\n Connection %s\n", s.config.getDbUrl())
	}

	s.db = db

	s.initRepo()

	return nil
}

func (s *Store) initRepo() {
	s.Users = usersrepository.New(s.db)
}
