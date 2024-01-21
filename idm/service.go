package idm

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrNoPassword = errors.New("idm: No password found")

type Service struct {
	queries *Queries
}

func (s *Service) FindPasswwordHash(username string) (string, error) {
	return "", nil
}

func (s *Service) VerifyPassword(username string, password string) (bool, error) {
	hash, err := s.FindPasswwordHash(username)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
