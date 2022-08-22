package service

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type authService struct{}

func (service *authService) JwtGenerateToken(userID primitive.ObjectID) (string, error) {
	JWT_SECRET_KEY := []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	signedToken, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
