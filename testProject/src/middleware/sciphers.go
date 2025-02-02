package middleware

import (
	"errors"
	"os"
)

func GetEnvSecret() (string, error) {
	key := os.Getenv("SECRET_KEY")

	if key == "" {
		return "", errors.New("SECRET_KEY not set")
	}

	return key, nil
}
