package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) Save() (*User, error) {
	if err := DB.Create(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(_ *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func FindByUsername(username string) (*User, error) {
	u := User{}
	if err := DB.Model(User{}).Where("username = ?", username).Take(&u).Error; err != nil {
		return &u, err
	}
	return &u, nil
}

func FindById(id uint) (User, error) {
	u := User{}
	if err := DB.First(&u, id).Error; err != nil {
		return u, errors.New("user not found")
	}

	u.stripPassword()

	return u, nil
}

func (u *User) stripPassword() {
	u.Password = ""
}
