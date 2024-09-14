package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	Password := string(bytes)
	return Password, nil
}

func CheckPassword(providedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	return err == nil
}