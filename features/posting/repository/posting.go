package repository

import (
	"sosmed/features/posting"
	model "sosmed/features/users/repository"


	"gorm.io/gorm"
)

type PostingModel struct {
	gorm.Model
	Postingan string
	Foto      string
	UserID    uint
}

type postingQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) posting.Repo {
	return &postingQuery{
		db: db,
	}
}

func (ip *postingQuery) InsertPosting(userID uint, newPosting posting.Posting) (posting.Posting, error) {

	var inputData = new(PostingModel)
	inputData.UserID = userID
	inputData.Postingan = newPosting.Postingan
	inputData.Foto = newPosting.Foto

	if err := ip.db.Create(&inputData).Error; err != nil {
		return posting.Posting{}, err
	}
	var user model.UserModel
	ip.db.Table("user_models").Where("id = ?", userID).Find(&user)

	newPosting.Users = user

	var Post PostingModel
	ip.db.Table("posting_models").Where("user_id = ?", userID).Last(&Post)
	newPosting.ID = userID
	newPosting.ID = Post.ID
	newPosting.Foto = Post.Foto
	newPosting.Postingan = Post.Postingan

	return newPosting, nil
}
