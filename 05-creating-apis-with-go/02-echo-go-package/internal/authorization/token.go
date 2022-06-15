package authorization

import (
	"errors"
	"time"

	"go-api/internal/model"
	"github.com/dgrijalva/jwt-go"
)

// GenerateToken .
func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Albert",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken .
func ValidateToken(t string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claim{}, verifyFunction)
	if err != nil {
		return model.Claim{}, err
	}
	if !token.Valid {
		return model.Claim{}, errors.New("invalid token")
	}

	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("can't obtain the claim")
	}

	return *claim, nil
}

// This validation should be done on the storage directory
func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
