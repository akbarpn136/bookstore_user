package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusCreated, "Created new user")
}

func GetUser(c *gin.Context) {
	userId, _ := c.Params.Get("user_id")
	c.String(http.StatusOK, "Get user with id: "+userId)
}
