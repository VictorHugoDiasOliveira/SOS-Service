package request

// UserLogin represents the data required for user login
// @Summary User Login Data
// @Description Structure containing the necessary fields for user login
type UserLogin struct {
	// User's email (required and must be a valid email address)
	Email string `json:"email" binding:"required,email" example:"email@gmail.com"`

	// User's password (required and must be between 5 and 50 characters)
	Password string `json:"password" binding:"required,min=5,max=50" example:"senha1234"`
}
