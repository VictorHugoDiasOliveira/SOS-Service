package response

type UserResponse struct {
	ID    string `json:"id" example:"6688570733fbe35fc2f33a15"`
	Email string `json:"email" example:"email@gmail.com"`
	Name  string `json:"name" example:"Larissa"`
	Age   int    `json:"age" example:"22"`
}
