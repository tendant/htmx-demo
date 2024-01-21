package idm

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrNoPassword = errors.New("idm: No password found")

func FindPasswwordHash(email string) (string, error) {
	return "", nil
}

func VerifyPassword(email string, password string) (bool, error) {
	hash, err := FindPasswwordHash(email)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
