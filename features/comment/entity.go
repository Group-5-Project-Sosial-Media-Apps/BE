package comment

import (
	model "sosmed/features/users/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Comment struct {
	ID     uint
	Pesan  string
	PostID uint
	Users  model.UserModel
}

type Handler interface {
	Add() echo.HandlerFunc
	DelComment() echo.HandlerFunc
}

type Service interface {
	TambahComment(token *jwt.Token, newComment Comment) (Comment, error)
	DelComment(commentID uint) (Comment, error)
}

type Repo interface {
	InsertComment(userID uint, newComment Comment) (Comment, error)
	DelComment(commentID uint) (Comment, error)
}
