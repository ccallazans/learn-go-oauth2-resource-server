package info

import (
	"log"
	"net/http"

	"github.com/ccallazans/learn-go-oauth2-resource-server/internal/auth"
	"github.com/labstack/echo/v4"
)

type InfoHandler struct{}

func NewInfoHandler() *InfoHandler {
	return &InfoHandler{}
}

type getInfoResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	User    *auth.Payload `json:"user"`
}

func (h *InfoHandler) GetInfo(c echo.Context) error {

	userDetails, exists := c.Get("user").(*auth.Payload)
	if !exists {
		log.Println("Error getting user from context!")
	}

	return c.JSON(http.StatusOK, &getInfoResponse{
		Status:  "OK",
		Message: "This is a protected resource",
		User:    userDetails,
	})
}
