package comment

type CommentRequest struct {
	PostingID uint   `json:"postingid"`
	Pesan     string `json:"pesan"`
}

type CommentResponse struct {
	CommentID uint   `json:"Comment_id"`
	Pesan     string `json:"pesan"`
	User      CommentResponseUser
}

type CommentResponseUser struct {
	UserID   uint   `form:"user_id"`
	Nama     string `form:"nama"`
	UserName string `form:"username"`
	Foto     string `form:"foto"`
}
