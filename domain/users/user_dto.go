package users

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() []string {
	var messages []string
	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			messages = append(messages, "Kolom "+er.Field()+" harus memenuhi "+er.Tag())
		}

		return messages
	}

	return nil
}
