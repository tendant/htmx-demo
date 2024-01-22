package idm

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
)

var ErrNoPassword = errors.New("idm: No password found")

type Service struct {
	Queries *Queries
}

func (s *Service) FindPasswwordHash(username string) ([]byte, error) {
	row, err := s.Queries.FindUserByUsername(context.Background(), username)
	if err != nil {
		slog.Error("Failed FindUserByUsername", "username", username)
		return nil, err
	}
	return row.Password, nil
}

func (s *Service) VerifyPassword(username string, password string) (bool, error) {
	hash, err := s.FindPasswwordHash(username)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
