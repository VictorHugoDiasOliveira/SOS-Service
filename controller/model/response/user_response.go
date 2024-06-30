package response

type UserResponse struct {
	ID      uint   `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
}
