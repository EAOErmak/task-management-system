package handlers

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

type loginResponse struct {
	Token string `json:"token"`
}

type registerResponse struct {
	Token string      `json:"token"`
	User  authUserDTO `json:"user"`
}

type authUserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
