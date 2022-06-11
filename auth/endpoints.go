package auth

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type AuthResponse struct {
	Token string
}

type Endpoints struct {
	Auth endpoint.Endpoint
}

func MakeEndpoints(s AuthSvc) Endpoints {
	return Endpoints{
		Auth: makeAuthEndpoint(s),
	}
}

func makeAuthEndpoint(s AuthSvc) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		userId, err := bearerTokenFromContext(ctx)
		if err != nil {
			return AuthResponse{}, &APIError{Code: http.StatusUnauthorized, Err: err}
		}
		token, err := s.Auth(ctx, userId)
		if err != nil {
			e := &APIError{Err: err}
			if err == ErrUserNotFound {
				e.Code = http.StatusUnauthorized
			}
			return AuthResponse{}, e
		}
		return AuthResponse{Token: token}, nil
	}
}
