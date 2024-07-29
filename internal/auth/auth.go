package auth

import (
	"errors"
	"net/http"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header not found")
	}
	apiKey := authHeader[len("ApiKey "):]
	if len(apiKey) == 0 {
		return "", errors.New("no token")
	}
	return apiKey, nil
}
