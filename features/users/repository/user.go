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

	newUser.UserID = inputDB.ID

	return newUser, nil
}

func (uq *userQuery) Login(username string) (users.User, error) {
	var userData = new(UserModel)

	if err := uq.db.Where("user_name", username).First(userData).Error; err != nil {
		return users.User{}, err
	}

	var result = new(users.User)

	result.UserID = userData.ID
	result.Nama = userData.Nama
	result.Password = userData.Password
	result.UserName = userData.UserName
	
	return *result, nil
}

func (uq *userQuery) GetUserById(id uint) (users.User, error) {
	var userData = new(UserModel)

	if err := uq.db.Where("id", id).Find(&userData).Error; err != nil {
		return users.User{}, err
	}

	var result = new(users.User)

	result.UserID = userData.ID
	result.Nama = userData.Nama
	result.UserName = userData.UserName
	result.Email = userData.Email
	result.Foto = userData.Foto

	return *result, nil
}

func (uq *userQuery) DelUserById(id uint) (users.User, error) {
	var userData = new(UserModel)

	if err := uq.db.Where("id", id).Find(&userData).Error; err != nil {
		return users.User{}, err
	}

	var result = new(users.User)

	result.UserID = userData.ID
	result.Nama = userData.Nama
	result.UserName = userData.UserName
	result.Email = userData.Email

	uq.db.Where("id", id).Delete(&userData)
	return *result, nil
}


func (us *userQuery) UpdateUser(id uint, updateUser users.User) (users.User, error) {
	var existingUser = new(UserModel)
	existingUser.Nama = updateUser.Nama
	existingUser.UserName = updateUser.UserName
	existingUser.Foto = updateUser.Foto

	if err := us.db.Where("id = ?", id).Updates(existingUser).Error; err != nil {
		return users.User{}, err
	}


	if updateUser.UserID != 0 {
		existingUser.ID = updateUser.UserID

	}

	if updateUser.Nama != "" {
		existingUser.Nama = updateUser.Nama
	}

	if updateUser.UserName != "" {
		existingUser.UserName = updateUser.UserName
	}

	if updateUser.Foto != "" {
		existingUser.Foto = updateUser.Foto
	}




	result := users.User{

		UserID: id,

		Nama: existingUser.Nama,
		UserName: existingUser.UserName,
		Foto: existingUser.Foto,
	}

	return result, nil


}