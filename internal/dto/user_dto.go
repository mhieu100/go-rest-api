package dto

type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required,min=3,max=100"`
}

type UpdateUserRequest struct {
	Email string `json:"email" binding:"omitempty,email"`
	Name  string `json:"name" binding:"omitempty,min=3,max=100"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,name"`
	Password string `json:"password" binding:"required,password,min=6"`
}

type RegisterResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
