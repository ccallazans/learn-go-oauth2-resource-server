package api

import (
	"os"

	"github.com/ccallazans/learn-go-oauth2-resource-server/internal/api/middlewares"
	"github.com/ccallazans/learn-go-oauth2-resource-server/internal/api/v1/info"
	"github.com/ccallazans/learn-go-oauth2-resource-server/internal/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	// Router
	r := echo.New()

	// Default Middlewares
	r.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			}),
		middleware.LoggerWithConfig(
			middleware.LoggerConfig{
				Format: "method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
			}),
	)

	// Handlers
	infoHandler := info.NewInfoHandler()

	// Middlewares
	googleAuth := auth.NewGoogleOAuth2Provider(os.Getenv("CLIENT_ID"))
	authMiddleware := middlewares.NewAuthMiddleware(googleAuth).Auth

	// Endpoints
	actuator := r.Group("/actuator")
	actuator.GET("/health", Health)

	v1 := r.Group("/v1")
	v1.Use(authMiddleware)
	v1.GET("/info", infoHandler.GetInfo)

	return r
}
