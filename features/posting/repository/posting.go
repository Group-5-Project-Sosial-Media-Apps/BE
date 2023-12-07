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
	User      model.UserModel `gorm:"foreignKey:UserID"`
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

	if err := ga.db.Offset(offset).Limit(pageSize).Preload("User").Find(&postings).Error; err != nil {
		return nil, 0, err
	}

	var result []posting.Posting
	for _, s := range postings {
		result = append(result, posting.Posting{
			ID:        s.ID,
			Postingan: s.Postingan,
			Foto:      s.Foto,
			Users:     s.User,
		
		})
	}

	return result, int(totalCount), nil
}
