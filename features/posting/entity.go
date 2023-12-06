package posting

import (
	model "sosmed/features/users/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Posting struct {
	ID        uint
	Postingan string
	Foto      string
	Users     model.UserModel
}

type Handler interface {
	Add() echo.HandlerFunc
}

type Service interface {
	TambahPosting(token *jwt.Token, newPosting Posting) (Posting, error)
}

type Repo interface {
	InsertPosting(userID uint, newPosting Posting) (Posting, error)
}
