package user

type RegisterUserInput struct {
	Name        string `json:"name" binding:"required"`
	NoHandphone string `json:"no_handphone" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type UserSearchInput struct {
	Column string `json:"column" binding:"required"`
	Value  string `json:"value" binding:"required"`
}
