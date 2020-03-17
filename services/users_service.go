package services

import (
	"bookstore/domain/users"
	"bookstore/utils/errors"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func CreateUser(user users.User) (*users.User, []errors.ValidationErrors) {
	var messages []errors.ValidationErrors

	if err := validate.Struct(&user); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			messages = append(messages, errors.ValidationErrors{
				Field: e.Field(),
				Tag:   e.Tag(),
			})
		}

		return nil, messages
	}

	return &user, nil
}
