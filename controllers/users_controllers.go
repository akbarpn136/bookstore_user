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
		c.JSON(http.StatusBadRequest, errors.BadRequest([]string{err.Error()}))

		return
	}

	user := users.User{
		Id:          val.Get("id"),
		FirstName:   val.Get("first_name"),
		LastName:    val.Get("last_name"),
		Email:       val.Get("email"),
		DateCreated: val.Get("date_created"),
	}

	result, Err := services.CreateUser(user)

	if Err != nil {
		c.JSON(http.StatusBadRequest, Err)

		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, _ := c.Params.Get("user_id")
	result, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(http.StatusNotFound, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
