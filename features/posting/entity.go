package posting

import (
	"sosmed/features/users"
	model "sosmed/features/users/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Posting struct {
	ID        uint `json:"id"`
	Postingan string
	Foto      string
	Users     model.UserModel
	User      []users.User
}

type Handler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

type Service interface {
	TambahPosting(token *jwt.Token, newPosting Posting) (Posting, error)
	GetAllPosting(page, pageSize int) ([]Posting, int, error)
}

type Repo interface {
	InsertPosting(userID uint, newPosting Posting) (Posting, error)
	GetAllPosting(page, pageSize int) ([]Posting, int, error)
}
