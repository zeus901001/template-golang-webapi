package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique"`
	Password     string
	FirstName    string
	LastName     string
	Role         int8
	Permission   bool
	RefreshToken string
}

/* Hash password. */
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil
	}
	user.Password = string(bytes)
	return nil
}

/* Verify hashed password. */
func (user *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
