package model

import (
	"sosmed/features/users"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Nama     string
	UserName string
	Password string
	Foto     string
	Email    string
}

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.Repository {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Register(newUser users.User) (users.User, error) {
	var inputDB = new(UserModel)
	inputDB.Nama = newUser.Nama
	inputDB.Email = newUser.Email
	inputDB.UserName = newUser.UserName
	inputDB.Password = newUser.Password

	if err := uq.db.Create(&inputDB).Error; err != nil {
		return users.User{}, err
	}

	newUser.ID = inputDB.ID

	return newUser, nil
}

func (uq *userQuery) Login(username string) (users.User, error) {
	var userData = new(UserModel)

	if err := uq.db.Where("user_name", username).First(userData).Error; err != nil {
		return users.User{}, err
	}

	var result = new(users.User)

	result.ID = userData.ID
	result.Nama = userData.Nama
	result.Password = userData.Password

	return *result, nil
}
