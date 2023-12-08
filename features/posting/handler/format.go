package posting

type PostingRequest struct {
	Pesan string `form:"pesan"`
	Foto  string `form:"foto"`
}

type PostingResponse struct {
	PostingID uint   `json:"posting_id"`
	Pesan     string `json:"pesan"`
	Foto      string `json:"foto"`
	// UserID    uint   `json:"userid"`
	// UserName  string `json:"username"`
	User    PostingResponseUser `json:"user"`
	Comment []CommentResponse   `json:"comment"`
}

type CommentResponse struct {
	CommentID uint                `json:"Comment_id"`
	Pesan     string              `json:"pesan"`
	User      PostingResponseUser `json:"user"`
}

type PostingResponseUser struct {
	UserID   uint   `json:"user_id"`
	Nama     string `json:"nama"`
	UserName string `json:"username"`
	Foto     string `json:"foto"`
}

type PostingUpdate struct {
	PostingID uint `form:"postingid"`
	Posting string `form:"posting"`
	Foto     string `form:"foto"`
}