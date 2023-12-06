package posting

type PostingRequest struct {
	Pesan string `form:"pesan"`
	Foto  string `form:"foto"`
}

type PostingResponse struct {
	PostingID uint   `json:"posting_id"`
	Pesan     string `json:"pesan"`
	Foto      string `json:"foto"`
	User      PostingResponseUser
}

type PostingResponseUser struct {
	UserID   uint   `form:"user_id"`
	Nama     string `form:"nama"`
	UserName string `form:"username"`
	Foto     string `form:"foto"`
}
