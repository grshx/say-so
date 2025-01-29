package handlers

import (
	"errors"

	"github.com/grshx/say-so/user-service/models"
)

func validateCredentials(user models.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("username or password cannot be empty")
	}
	return nil
}

func RegisterUser(newUser models.User) error {

	err := validateCredentials(newUser)

	if err != nil {
		return err
	}

	return nil
}

func Login(user models.User) (string, error) {

	err := validateCredentials(user)

	if err != nil {
		return "", err
	}

	return "", nil
}
