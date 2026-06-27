package users

import (
	"fmt"

	"github.com/noboKumar/SpotSync-server/auth"
	"github.com/noboKumar/SpotSync-server/domain/users/dto"
)

var ErrInvalidCredentials = fmt.Errorf("invalid email or password")

type service struct {
	repo       Repository
	jwtService auth.JwtService
}

func NewService(repo Repository, jwtService auth.JwtService) *service {
	return &service{repo: repo, jwtService: jwtService}
}

func (s *service) CreateUser(req dto.CreateRequest) (*dto.Response, error) {
	// Create a new user entity from the request data
	user := User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	// Set the password for the user (this will hash the password)
	if err := user.SetPassword(req.Password); err != nil {
		return nil, err
	}

	// Set the role for the user; default to "driver" if not provided
	if req.Role == "" {
		req.Role = "driver"
	}

	if err := s.repo.CreateUser(&user); err != nil {
		return nil, err
	}

	return &dto.Response{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
