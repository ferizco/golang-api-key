package main

import (
	"errors"
	"regexp"
)

var validAPIKeys = map[string]string{
	"key-123": "admin",
	"key-456": "user",
}

func ValidateAPIKey(apikey string) (string, bool) {
	role, exists := validAPIKeys[apikey]
	return role, exists
}

func ValidateLogin(username, password string) error {
	if username == "" {
		return errors.New("username tidak boleh kosong")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username) {
		return errors.New("username hanya boleh mengandung huruf dan angka")
	}

	if len(password) < 8 {
		return errors.New("password harus memiliki minimal 8 karakter")
	}

	return nil
}
