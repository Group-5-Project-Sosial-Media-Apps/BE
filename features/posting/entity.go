package posting

import (
	"sosmed/features/comment"
	"sosmed/features/users"
	model "sosmed/features/users/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Posting struct {
	ID        uint `json:"id"`
	Postingan string
	Foto      string
	UserID    uint
	Users     model.UserModel
	User      []users.User
	Comment   []comment.Comment
}

type Handler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	DelPost() echo.HandlerFunc
}

type Service interface {
	TambahPosting(token *jwt.Token, newPosting Posting) (Posting, error)
	GetAllPosting(page, pageSize int) ([]Posting, int, error)
	GetPostingById(userID uint) ([]Posting, error)
	DelPost(PostID uint) (Posting, error)
}

type Repo interface {
	InsertPosting(userID uint, newPosting Posting) (Posting, error)
	GetAllPosting(page, pageSize int) ([]Posting, int, error)
	GetPostingById(userID uint) ([]Posting, error)
	DelPost(PostID uint) (Posting, error)
}
