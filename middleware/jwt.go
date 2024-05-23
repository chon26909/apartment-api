package middleware

import (
	"strings"

	"github.com/chon26909/apartment-api/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func VerifyHeader() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.ErrUnauthorized
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 {
				return echo.ErrUnauthorized
			}

			tokenStr := bearerToken[1]

			jwtSecret := []byte("abcd")

			claims := &models.CustomClaims{}

			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})

			if err != nil || !token.Valid {
				return echo.ErrUnauthorized
			}

			c.Set("userId", claims.UserId)
			return next(c)
		}
	}
}
