package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5,max=50"`
	Name     string `json:"name" binding:"required,min=5,max=100"`
	Age      int    `json:"age" binding:"required,min=18"`
}

type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=5,max=50"`
	Age  int    `json:"age" binding:"omitempty,min=18"`
}
