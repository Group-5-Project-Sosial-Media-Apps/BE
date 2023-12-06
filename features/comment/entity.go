package comment

import (
	model "sosmed/features/users/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Comment struct {
	ID    uint
	Pesan  string
	Users model.UserModel
}

type Handler interface {
	Add() echo.HandlerFunc
}

type Service interface {
	TambahComment(token *jwt.Token, newComment Comment) (Comment, error)
}

type Repo interface {
	InsertComment(userID uint, newComment Comment) (Comment, error)
}
