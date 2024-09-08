package rest

import (
	"context"
	"net/http"
	"strings"

	"github.com/ReyLegar/ShoppingList/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userContentKey = contextKey("user")

var JwtKey = []byte("pass_secret_one")

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получение заголовка Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Отсутствует Authorization заголовок", http.StatusUnauthorized)
			return
		}

		// Ожидаем формат: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Неверный формат Authorization заголовка", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]

		// Парсинг и проверка токена
		claims := &models.Claims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil || !tkn.Valid {
			http.Error(w, "Неверный или просроченный токен", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userContentKey, claims.Login)
		// Передача управления дальше
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
