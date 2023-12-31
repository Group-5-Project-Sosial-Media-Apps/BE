package repository

import (
	"sosmed/features/comment"
	model "sosmed/features/users/repository"

	"gorm.io/gorm"
)

type CommentModel struct {
	gorm.Model
	Pesan  string
	UserID uint
	PostId uint
	User   model.UserModel `gorm:"foreignKey:UserID"`
}

type commentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.Repo {
	return &commentQuery{
		db: db,
	}
}

func (ic *commentQuery) InsertComment(userID uint, newComment comment.Comment) (comment.Comment, error) {

	var inputData = new(CommentModel)
	inputData.UserID = userID
	inputData.Pesan = newComment.Pesan
	inputData.PostId = newComment.PostID

	if err := ic.db.Create(&inputData).Error; err != nil {
		return comment.Comment{}, err
	}
	var user model.UserModel
	ic.db.Table("user_models").Where("id = ?", userID).Find(&user)

	newComment.Users = user

	var Post CommentModel
	ic.db.Table("comment_models").Where("user_id = ?", userID).Last(&Post)
	newComment.ID = userID
	newComment.ID = Post.ID
	newComment.Pesan = Post.Pesan

	return newComment, nil
}

func (uq *commentQuery) DelComment(CommentID uint) (comment.Comment, error) {
	var commentData = new(CommentModel)

	if err := uq.db.Where("id", CommentID).Find(&commentData).Error; err != nil {
		return comment.Comment{}, err
	}

	var result = new(comment.Comment)
	result.ID = commentData.ID

	uq.db.Where("id", CommentID).Delete(&commentData)
	return *result, nil
}
