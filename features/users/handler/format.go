package user

type UserRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Nama     string `json:"nama"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Nama     string `json:"nama"`
	UserName string `json:"username"`
	Token    string `json:"token"`
}
