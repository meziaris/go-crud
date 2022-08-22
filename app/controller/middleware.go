package controller

import (
	"context"
	"errors"
	"go-crud/app/helper"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		tokenHeader := strings.Split(header, " ")

		if len(tokenHeader) < 2 {
			helper.WebResponse(w, http.StatusUnauthorized, "ERROR", errors.New("Unauthorization"))
			return
		}

		tokenString := tokenHeader[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				helper.WebResponse(w, http.StatusUnauthorized, "ERROR", errors.New("Unauthorization"))
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			helper.WebResponse(w, http.StatusUnauthorized, "ERROR", errors.New("Unauthorization"))
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			helper.WebResponse(w, http.StatusUnauthorized, "ERROR", errors.New("Unauthorization"))
			return
		}

		userID := claims["user_id"]
		const key helper.KeyType = "currentUserID"

		ctx := context.WithValue(r.Context(), key, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
