package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
)

type Claims struct {
	jwt.StandardClaims
}

func (au *authRepo) CreateUserSession(userUniqueID string) (models.UserSession, error) {
	accessToken, err := au.generateAccessToken(userUniqueID)
	if err != nil {
		return models.UserSession{}, err
	}

	return models.UserSession{
		JWTToken: accessToken,
	}, nil
}

func (au *authRepo) generateAccessToken(userUniqueID string) (string, error) {
	expTime := time.Now().Add(au.expTime).Unix() // Token berlaku selama 24 jam
	accessClaims := Claims {
		jwt.StandardClaims{
			Subject: userUniqueID,
			ExpiresAt: expTime,
		},
	}
	accessJWT := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims)

	return accessJWT.SignedString(au.signKey)
}

func (au *authRepo) CheckSession(data models.UserSession) (userUniqueID string, err error) {
	accessToken, err := jwt.ParseWithClaims(data.JWTToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return &au.signKey.PublicKey, nil // pakai pointer
	})
	fmt.Println("DEBUG: JWTToken =", data.JWTToken)
	fmt.Printf("DEBUG: signKey.PublicKey type = %T\n", au.signKey.PublicKey)
	if err != nil {
		return " ", err
	}

	acceessTokenClaims, ok := accessToken.Claims.(*Claims)
	if !ok || !accessToken.Valid {
		return " ", jwt.NewValidationError("invalid token, unauthorized", jwt.ValidationErrorMalformed)
	}
	fmt.Println("DEBUG: JWTToken =", data.JWTToken)
	if accessToken.Valid {
		return acceessTokenClaims.Subject, nil
	}

	return " ", jwt.NewValidationError("invalid token, unauthorized", jwt.ValidationErrorMalformed)
}
