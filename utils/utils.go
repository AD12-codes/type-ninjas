package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashedPassword(password string) ([]byte, error) {
	hashedPassword, generatePasswordError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if generatePasswordError != nil {
		return nil, generatePasswordError
	}

	return hashedPassword, nil
}
