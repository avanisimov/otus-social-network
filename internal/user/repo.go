package user

import (
	"context"
	"fmt"
)

type Repository struct {
	Users map[string]*User
}

func NewRepository() *Repository {
	return &Repository{
		Users: make(map[string]*User),
	}
}

func (r *Repository) GetUserByID(ctx context.Context, userID string) (*User, error) {
	return r.Users[userID], nil
}

func (r *Repository) CreateUser(user User) (string, error) {
	id := fmt.Sprintf("user-%d", len(r.Users))
	user.ID = id
	r.Users[id] = &user
	return id, nil
}
