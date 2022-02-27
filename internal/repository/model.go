package repository

import (
	"errors"
	"gorm.io/gorm"
)

const AllowedLength = 6

type User struct {
	gorm.Model
	Username string
	Password string
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		return errors.New("you cannot delete the boss lah")
	}
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		return errors.New("you cannot update the boss lah")
	}
	if !u.validatePassword() {
		return errors.New("password mush be at-least 6 characters")
	}
	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if !u.validatePassword() {
		return errors.New("password mush be at-least 6 characters")
	}
	return
}

func (u *User) validatePassword() bool {
	return !(len(u.Password) < AllowedLength)
}
