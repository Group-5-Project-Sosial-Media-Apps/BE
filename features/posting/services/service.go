package services

import (
	"errors"
	"sosmed/features/posting"
	"sosmed/helper/jwt"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type PostingServices struct {
	m posting.Repo
}

func New(model posting.Repo) posting.Service {
	return &PostingServices{
		m: model,
	}
}

func (tp *PostingServices) TambahPosting(token *golangjwt.Token, newPosting posting.Posting) (posting.Posting, error) {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return posting.Posting{}, err
	}

	result, err := tp.m.InsertPosting(userID, newPosting)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return posting.Posting{}, errors.New("barang sudah pernah diinputkan")
		}
		return posting.Posting{}, errors.New("terjadi kesalahan pada server")
	}

	return result, nil
}

func (ga *PostingServices) GetAllPosting(page int, pageSize int) ([]posting.Posting, int, error) {
	postings, totalCount, err := ga.m.GetAllPosting(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return postings, totalCount, nil
}
