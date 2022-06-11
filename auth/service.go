package auth

import (
	"context"
)

type AuthSvc interface {
	Auth(ctx context.Context, userId string) (token string, err error)
}

func NewService(repo Repository, tokenBuilder TokenBuilder) AuthSvc {
	return &authSvc{
		repo:    repo,
		builder: tokenBuilder,
	}
}

type authSvc struct {
	repo    Repository
	builder TokenBuilder
}

func (s *authSvc) Auth(ctx context.Context, userId string) (token string, err error) {
	user, err := s.repo.GetUser(ctx, userId)
	if err != nil {
		return "", err
	}

	token, err = s.builder.Build(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
