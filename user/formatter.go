package user

import "dbo-be/entities"

type UserFormatter struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	NoHandphone string `json:"no_handphone"`
	Email       string `json:"email"`
	Token       string `json:"token"`
}

func FormatUser(user entities.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		Name:        user.Name,
		NoHandphone: user.NoHandphone,
		Email:       user.Email,
		Token:       token,
	}

	return formatter
}
