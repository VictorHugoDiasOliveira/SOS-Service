package response

type UserResponse struct {
	Id      uint   `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
}
