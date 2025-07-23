package utils

import (
	"net/http"
	"strings"
	"errors"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
)

func GetSessionData(r *http.Request) (models.UserSession, error) {
	authString := r.Header.Get("Authorization")
	if authString == "" {
		return models.UserSession{}, errors.New("authorization header missing")
	}
	if !strings.HasPrefix(authString, "Bearer ") {
		return models.UserSession{}, errors.New("invalid authorization header format")
	}
	splitString := strings.SplitN(authString, " ", 2)
	if len(splitString) != 2 || splitString[1] == "" {
		return models.UserSession{}, errors.New("invalid authorization header")
	}

	accessString := splitString[1]

	return models.UserSession{
		JWTToken: accessString,
	}, nil
}