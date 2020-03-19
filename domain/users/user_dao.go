package users

import (
	"bookstore/datasources/mysql/users_db"
	"bookstore/utils/date_utils"
	"bookstore/utils/errors"
)

var (
	usersDB = make(map[string]*User)
)

func (user *User) Get() *errors.RestErrors {
	result := usersDB[user.Id]

	if result == nil {
		return errors.NotFound([]string{"User not found"})
	}

	if err := users_db.Db.Ping(); err != nil {
		panic(err)
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
		return errors.NotAcceptable([]string{"User already exist"})
	}

	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user

	return nil
}
