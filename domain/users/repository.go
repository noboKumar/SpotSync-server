package users

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {

	// create a new user in the database
	CreateUser(user *User) error

	// get a user by email from the database
	GetUserByEmail(email string) (*User, error)
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

func (r *repository) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
