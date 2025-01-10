package service

import (
	"database/sql"

	"github.com/A-Victory/user-mig/user/db"
	"github.com/A-Victory/user-mig/user/models"
)

type Service struct {
	db *db.DBconn
}

func NewServiceConn(db *db.DBconn) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateNewUser(user models.User) error {

	err := s.db.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUserByEmail(email string) (models.User, error) {
	user, err := s.db.GetUser(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, err
	}

	return *user, err
}

func (s *Service) ListAllUsers() ([]models.User, error) {

	users, err := s.db.ListUsers()
	if err != nil {
		return nil, err
	}

	return users, err
}
