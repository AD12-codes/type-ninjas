package users

import (
	"context"

	"github.com/AD12-codes/go-template/db"
	"github.com/AD12-codes/go-template/utils"
	"github.com/google/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, req CreateUserRequest) error
	GetAllUsers(ctx context.Context) ([]db.User, error)
}

type service struct {
	queries *db.Queries
}

func NewService(q *db.Queries) Service {
	return &service{queries: q}
}

// CreateUser implements Service.
func (s *service) CreateUser(ctx context.Context, req CreateUserRequest) error {

	id, uuidErr := uuid.NewV7()
	if uuidErr != nil {
		return uuidErr
	}

	hashedPassword, generatePasswordError := utils.GenerateHashedPassword(req.Password)
	if generatePasswordError != nil {
		return generatePasswordError
	}

	user := db.CreateUserParams{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  string(hashedPassword),
		Email:     req.Email,
		Username:  req.Username,
	}

	createUserDbError := s.queries.CreateUser(ctx, user)
	if createUserDbError != nil {
		return createUserDbError
	}
	return nil
}

func (s *service) GetAllUsers(ctx context.Context) ([]db.User, error) {
	users, getAllUsersError := s.queries.GetAllUsers(ctx)

	if getAllUsersError != nil {
		return nil, getAllUsersError
	}

	return users, nil
}
