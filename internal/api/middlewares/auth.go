package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/ccallazans/learn-go-oauth2-resource-server/internal/auth"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	Providers map[string]auth.OAuth2Provider
}

func NewAuthMiddleware(provider auth.OAuth2Provider) *AuthMiddleware {
	providers := map[string]auth.OAuth2Provider{
		"https://accounts.google.com": provider,
	}

	return &AuthMiddleware{
		Providers: providers,
	}
}

func (m *AuthMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
		}

		issuer, err := extractIssuer(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}

		provider, exists := m.Providers[issuer]
		if !exists {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid issuer")
		}

		user, err := provider.ValidateToken(token)
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("user", user)

		return next(c)
	}
}

func extractIssuer(token string) (string, error) {
	segments := strings.Split(token, ".")
	if len(segments) != 3 {
		return "", errors.New("invalid token format")
	}

	decoded, err := base64.RawURLEncoding.DecodeString(segments[1])
	if err != nil {
		return "", err
	}

	var payload struct {
		Issuer string `json:"iss"`
	}
	if err := json.Unmarshal(decoded, &payload); err != nil {
		return "", err
	}

	return payload.Issuer, nil
}
