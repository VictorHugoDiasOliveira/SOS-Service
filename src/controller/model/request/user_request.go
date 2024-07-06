package request

// UserRequest represents the input data for creating a new user
// @Summary User Input Data
// @Description Structure containing the required fields for creating a new user
type UserRequest struct {
	// User's email (required, email)
	// @json
	// @jsonTag email
	// @jsonExample email@gmail.com
	// @binding required,email
	Email string `json:"email" binding:"required,email" example:"email@gmail.com"`

	// User's password (required, minimum of 5 characters, maximum of 50 characters)
	// @json
	// @jsonTag password
	// @jsonExample senha1234
	// @binding required,min=5,max=50
	Password string `json:"password" binding:"required,min=5,max=50" example:"senha1234"`

	// User's name (required, minimum of 5 characters, maximum of 50 characters)
	// Example: Victor Hugo
	// @json
	// @jsonTag name
	// @jsonExample Victor Hugo
	// @binding required,min=5,max=50
	Name string `json:"name" binding:"required,min=5,max=50" example:"Victor Hugo"`

	// User's age (required, minimum of 18)
	// @json
	// @jsonTag age
	// @jsonExample 22
	// @binding required,min=5
	Age int `json:"age" binding:"required,min=18" example:"22"`
}

type UserUpdateRequest struct {
	// User's name (omitempty, minimum of 5 characters, maximum of 50 characters)
	// Example: Victor Hugo
	// @json
	// @jsonTag name
	// @jsonExample Victor Hugo
	// @binding omitempty,min=5,max=50
	Name string `json:"name" binding:"omitempty,min=5,max=50" example:"Victor Hugo"`

	// User's age (omitempty, minimum of 18)
	// @json
	// @jsonTag age
	// @jsonExample 22
	// @binding omitempty,min=5
	Age int `json:"age" binding:"omitempty,min=18" example:"22"`
}
