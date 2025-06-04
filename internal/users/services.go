package users

import (
	"context"

	"github.com/AD12-codes/type-ninjas/db"
	"github.com/AD12-codes/type-ninjas/utils"
	"github.com/google/uuid"
)

type Service interface {
	RegisterUser(c context.Context, req RegisterUserRequest) error
	GetAllUsers(c context.Context) ([]db.User, error)
	GetUser(c context.Context, id uuid.UUID) (db.User, error)
}

type service struct {
	queries *db.Queries
}

func NewService(queries *db.Queries) Service {
	return &service{queries: queries}
}

func (s *service) RegisterUser(c context.Context, req RegisterUserRequest) error {
	id, _ := uuid.NewV7()
	username := utils.GenerateUsername()

	registerUserError := s.queries.RegisterUser(c, db.RegisterUserParams{
		ID:        id,
		Auth0ID:   req.Auth0ID,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  username,
	})
	if registerUserError != nil {
		return registerUserError
	}
	return nil
}

func (s *service) GetAllUsers(c context.Context) ([]db.User, error) {
	return s.queries.GetAllUsers(c)
}

func (s *service) GetUser(c context.Context, id uuid.UUID) (db.User, error) {
	return s.queries.GetUserById(c, id)
}
