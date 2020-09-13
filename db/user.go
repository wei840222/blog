package db

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string
	AvatarURL  sql.NullString
	LINEUserID sql.NullString `gorm:"size:33,unique"`
}

func CreateUser(user *User) (*User, error) {
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(id uint) (*User, error) {
	var user User
	if err := db.Where(User{
		Model: gorm.Model{
			ID: id,
		},
	}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
