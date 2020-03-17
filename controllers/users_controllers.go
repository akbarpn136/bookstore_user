package controllers

import (
	"bookstore/domain/users"
	"bookstore/services"
	"bookstore/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	err := c.Request.ParseForm()
	val := c.Request.Form

	if err != nil {
		c.JSON(http.StatusBadRequest, errors.RestErrors{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Error:   "bad_request",
		})
	}

	user := users.User{
		Id:          1,
		FirstName:   val.Get("first_name"),
		LastName:    val.Get("last_name"),
		Email:       val.Get("email"),
		DateCreated: val.Get("date_created"),
	}

	result, Err := services.CreateUser(user)

	if Err != nil {
		c.JSON(http.StatusBadRequest, errors.RestErrors{
			Message: Err.Error(),
			Code:    http.StatusBadRequest,
			Error:   "bad_request",
		})
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, _ := c.Params.Get("user_id")
	c.String(http.StatusOK, "Get user with id: "+userId)
}
