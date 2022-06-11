package auth

import (
	"context"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
}

type Repository interface {
	GetUser(ctx context.Context, userID string) (User, error)
}

func NewRepository() Repository {
	return &inMemRepository{
		users: map[string]User{
			"1122334455": {
				ID:        "1122334455",
				FirstName: "John",
				LastName:  "Doe",
				Username:  "jdoe",
				Email:     "john.doe@example.com",
			},
			"1122334466": {
				ID:        "1122334466",
				FirstName: "Jane",
				LastName:  "Doe",
				Username:  "jadoe",
				Email:     "jane.doe@example.com",
			},
		},
	}
}

type inMemRepository struct {
	users map[string]User
}

func (r *inMemRepository) GetUser(ctx context.Context, userID string) (User, error) {
	user, ok := r.users[userID]
	if ok {
		return user, nil
	}
	return User{}, ErrUserNotFound
}
