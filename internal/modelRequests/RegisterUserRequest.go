package modelRequests


type RegisterUserRequests struct {
	Login    string `json:"login" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Email    string  `json:"email" binding:"required"`
}

