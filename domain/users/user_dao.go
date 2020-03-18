package users

import (
	"bookstore/utils/errors"
	"net/http"
)

var (
	usersDB = make(map[string]*User)
)

func (user *User) Get() *errors.RestErrors {
	result := usersDB[user.Id]

	if result == nil {
		return &errors.RestErrors{
			Message: []string{"User not found"},
			Code:    http.StatusNotFound,
			Error:   "not_found",
		}
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErrors {
	currentUser := usersDB[user.Id]

	if currentUser != nil {
		return &errors.RestErrors{
			Message: []string{"User already exist"},
			Code:    http.StatusNotAcceptable,
			Error:   "not_acceptable",
		}
	}

	usersDB[user.Id] = user

	return nil
}
