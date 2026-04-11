package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error) {
	user := User{
		FirstName:    req.FirstName,
		SecondName:   req.SecondName,
		Birthdate:    req.Birthdate,
		Biography:    req.Biography,
		City:         req.City,
		PasswordHash: hashPassword(req.Password),
	}
	userID, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &RegisterResponse{UserID: userID}, nil
}

func hashPassword(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}
