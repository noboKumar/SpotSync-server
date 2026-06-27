package users

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {

	// create a new user in the database
	CreateUser(user *User) error
}

// repository is a concrete implementation of the Repository interface
type repository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of the user repository
func NewUserRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(user *User) error {
	// Use GORM to create a new user in the database
	result := r.db.Create(user)

	// Check for duplicate email error
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return errors.New("user with this email already exists")
		}
		return result.Error
	}
	return nil
}
