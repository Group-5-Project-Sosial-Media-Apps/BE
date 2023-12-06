package user

type UserRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	// UserID       uint   `json:"user_id"`
	Nama     string `json:"nama"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID       uint   `json:"user_id"`
	Nama     string `json:"nama"`
	UserName string `json:"username"`
	Token    string `json:"token"`
}

type GetUserByIdRequest struct {
	UserID uint `json:"user_id"`
}

type GetUserByIdResponse struct {
	UserID       uint   `json:"user_id"`
	Nama     string `json:"nama"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Foto     string `json:"foto"`
}

type DelUserByIdRequest struct {
	UserID uint `json:"user_id"`
}

type DelUserByIdResponse struct {
	UserID       uint   `json:"user_id"`
	Nama     string `json:"nama"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdate struct {
	UserID       uint   `form:"user_id"`
	Nama     string `form:"nama"`
	UserName string `form:"username"`
	Foto     string `form:"foto"`
}
