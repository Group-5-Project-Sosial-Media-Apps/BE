package services

import (
	"errors"
	"sosmed/features/comment"
	"sosmed/helper/jwt"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type CommentServices struct {
	m comment.Repo
}

func New(model comment.Repo) comment.Service {
	return &CommentServices{
		m: model,
	}
}

func (tc *CommentServices) TambahComment(token *golangjwt.Token, newComment comment.Comment) (comment.Comment, error) {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return comment.Comment{}, err
	}

	result, err := tc.m.InsertComment(userID, newComment)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return comment.Comment{}, errors.New("barang sudah pernah diinputkan")
		}
		return comment.Comment{}, errors.New("terjadi kesalahan pada server")
	}

	return result, nil
}

func (tc *CommentServices) DelComment(commentID uint) (comment.Comment, error) {
	result, err := tc.m.DelComment(commentID)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return comment.Comment{}, errors.New("username tidak ditemukan")
		}
		return comment.Comment{}, errors.New("terjadi kesalahan pada sistem")
	}
	return result, nil
}
