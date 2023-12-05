package services

import (
	"errors"
	"sosmed/features/users"
	"sosmed/helper/enkrip"
	"strings"
)

type userService struct {
	repo users.Repository
	h    enkrip.HashInterface
}

func New(r users.Repository, h enkrip.HashInterface) users.Service {
	return &userService{
		repo: r,
		h:    h,
	}
}

func (us *userService) Register(newUser users.User) (users.User, error) {
	// validasi
	if newUser.UserName == "" || newUser.Password == "" {
		return users.User{}, errors.New("incorrect input data")
	}

	// enkripsi password
	ePassword, err := us.h.HashPassword(newUser.Password)

	if err != nil {
		return users.User{}, errors.New("terdapat masalah saat memproses data")
	}

	newUser.Password = ePassword
	result, err := us.repo.Register(newUser)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return users.User{}, errors.New("data telah terdaftar pada sistem")
		}
		return users.User{}, errors.New("terjadi kesalahan pada sistem")
	}

	return result, nil
}

func (us *userService) Login(username string, password string) (users.User, error) {
	result, err := us.repo.Login(username)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return users.User{}, errors.New("username tidak ditemukan")
		}
		return users.User{}, errors.New("terjadi kesalahan pada sistem")
	}

	err = us.h.Compare(result.Password, password)

	if err != nil {
		return users.User{}, errors.New("password salah")
	}

	return result, nil
}

func (us *userService) GetUserById(id uint) (users.User, error) {
	result, err := us.repo.GetUserById(id)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return users.User{}, errors.New("username tidak ditemukan")
		}
		return users.User{}, errors.New("terjadi kesalahan pada sistem")
	}

	return result, nil
}

func (us *userService) DelUserById(id uint) (users.User, error) {
	result, err := us.repo.DelUserById(id)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return users.User{}, errors.New("username tidak ditemukan")
		}
		return users.User{}, errors.New("terjadi kesalahan pada sistem")
	}
	return result, nil
}
