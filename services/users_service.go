package services

import (
	"bookstore/domain/users"
	"bookstore/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErrors) {
	if err := user.Validate(); len(err) > 0 {
		return nil, errors.BadRequest(err)
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId string) (*users.User, *errors.RestErrors) {
	result := users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}
