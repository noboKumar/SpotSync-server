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

func (s *service) CreateUser(req dto.CreateRequest) (*dto.RegisterResponse, error) {
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

	return &dto.RegisterResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *service) LoginUser(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := user.CheckPassword(req.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, user.Role, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: accessToken,
		User: dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}
