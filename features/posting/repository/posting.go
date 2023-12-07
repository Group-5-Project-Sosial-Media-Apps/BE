package repository

import (
	"sosmed/features/comment"
	cr "sosmed/features/comment/repository"
	"sosmed/features/posting"
	model "sosmed/features/users/repository"

	"gorm.io/gorm"
)

type PostingModel struct {
	gorm.Model
	Postingan string
	Foto      string
	UserID    uint
	User      model.UserModel   `gorm:"foreignKey:UserID"`
	Comment   []cr.CommentModel `gorm:"foreignKey:PostId"`
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
	newPosting.ID = Post.ID
	newPosting.Foto = Post.Foto
	newPosting.Postingan = Post.Postingan

	return newPosting, nil
}

func (ga *postingQuery) GetAllPosting(page, pageSize int) ([]posting.Posting, int, error) {
	var postings []PostingModel
	var totalCount int64

	offset := (page - 1) * pageSize

	if err := ga.db.Model(&PostingModel{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := ga.db.Offset(offset).Limit(pageSize).Preload("User").Preload("Comment").Preload("Comment.User").Find(&postings).Error; err != nil {
		return nil, 0, err
	}

	var result []posting.Posting
	for _, s := range postings {
		tmp := new(posting.Posting)
		tmp.ID = s.ID
		tmp.Postingan = s.Postingan
		tmp.Foto = s.Foto
		tmp.Users = s.User

		for _, v := range s.Comment {
			tmp.Comment = append(tmp.Comment, comment.Comment{
				ID: v.ID,
				Pesan: v.Pesan,
				Users: v.User,
			})
		}

		result = append(result, *tmp)
		// result = append(result, posting.Posting{
		// 	ID:        s.ID,
		// 	Postingan: s.Postingan,
		// 	Foto:      s.Foto,
		// 	Users:     s.User,
		// })
	}

	return result, int(totalCount), nil
}
