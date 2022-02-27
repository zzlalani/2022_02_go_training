package repository

import "gorm.io/gorm"

type UserRepository interface {
	Insert(user *User) (uint, error)
	Update(user *User, id uint) error
	Delete(user *User) error
	Read(id uint) (*User, error)
}

type userImpl struct {
	db *gorm.DB
}

func (u *userImpl) Read(id uint) (*User, error) {
	var user User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userImpl) Insert(user *User) (uint, error) {
	errMsg := u.db.Create(&user).Error
	if errMsg != nil {
		return 0, errMsg
	}
	return user.ID, nil
}

func (u *userImpl) Update(user *User, id uint) (err error) {
	err = u.db.Model(user).Where("id = ?", id).Updates(&user).Error
	return
}

func (u *userImpl) Delete(user *User) error {
	return u.db.Delete(&user).Error
}


func NewUserRepo(db *gorm.DB) UserRepository {
	return &userImpl{
		db: db,
	}
}