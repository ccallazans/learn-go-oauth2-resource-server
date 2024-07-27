package auth

import (
	"context"

	"google.golang.org/api/idtoken"
)

// Interface
type OAuth2Provider interface {
	ValidateToken(token string) (*Payload, error)
}

// Google Implementation
type GoogleOAuth2Provider struct {
	ClientID string
}

func NewGoogleOAuth2Provider(clientID string) OAuth2Provider {
	return &GoogleOAuth2Provider{
		ClientID: clientID,
	}
}

func (o *GoogleOAuth2Provider) ValidateToken(token string) (*Payload, error) {
	payload, err := idtoken.Validate(context.Background(), token, o.ClientID)
	if err != nil {
		return nil, err
	}

	return &Payload{
		Issuer:   payload.Issuer,
		Audience: payload.Audience,
		Expires:  payload.Expires,
		IssuedAt: payload.IssuedAt,
		Subject:  payload.Subject,
		Claims:   payload.Claims,
	}, err
}
