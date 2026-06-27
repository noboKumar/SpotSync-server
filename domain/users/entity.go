package users

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id        uint      `gorm:"primaryKey"`
	name      string    `gorm:"not null"`
	email     string    `gorm:"unique;not null"`
	password  string    `gorm:"not null"`
	role      string    `gorm:"type:varchar(20);default:driver;not null"`
	createdAt time.Time `gorm:"autoCreateTime"`
	updatedAt time.Time `gorm:"autoUpdateTime"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err
}
