package Helpers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"unicode"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func VerifyPassword(password string) (valid bool) {
	/*
		Verify that password has 1 lowercase, 1 uppercase, 1 special character,
		1 digit and is at least 8 characters long
	*/

	if len(password) < 8 {
		return false
	}

	hasLowerCase := false
	hasUpperCase := false
	hasDigit := false
	hasSpecialCharacter := false

	for _, c := range password {
		if hasLowerCase && hasUpperCase && hasDigit && hasSpecialCharacter {
			return true
		}
		switch {
		case unicode.IsUpper(c):
			hasUpperCase = true
		case unicode.IsLower(c):
			hasLowerCase = true
		case unicode.IsDigit(c):
			hasDigit = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecialCharacter = true
		}
	}

	return hasLowerCase && hasUpperCase && hasDigit && hasSpecialCharacter
}
