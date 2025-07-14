package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("some_value")

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization Header Required", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token is not valid", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "username", claims.UserName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
