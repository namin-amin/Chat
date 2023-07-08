package models

// SignUpDto
//
// SignUp details required for getting creating new user
type SignUpDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	PassWord string `json:"password" validate:"required,gte=8"`
	Role     int    `json:"role"`
}

// SigInDto
//
// SignIn details required for logging in the user
type SigInDto struct {
	Email    string `json:"email" validate:"required,email"`
	PassWord string `json:"password" validate:"required,gte=8"`
}

// SignUpRespDto
//
// Response to signUp request if success
type SignUpRespDto struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
